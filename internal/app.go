// Package app
package app

// TODO: Refactor, split up, getting too big
import (
	"encoding/json"
	"fmt"
	Bootstrap "framework/app"
	appconfig "framework/app/config"
	commandexecutor "framework/internal/app/command"
	"framework/internal/app/config"
	"framework/internal/app/cron"
	"framework/internal/app/env"
	"framework/internal/app/request"
	"framework/internal/app/router"
	"framework/internal/app/session"
	"framework/internal/app/validator"
	internalconfig "framework/internal/internal-config"
	"net/http"
	"os"
	"reflect"
	"strconv"

	"github.com/olbrichattila/godi"
)

func New(container godi.Container) *App {
	app := &App{
		di: container,
		conf: config.New(
			appconfig.Routes,
			appconfig.Jobs,
			appconfig.Middlewares,
			appconfig.DiBindings,
			internalconfig.DiBindings,
			appconfig.ConsoleCommands,
			internalconfig.ConsoleCommands,
			appconfig.ViewFuncConfig,
			internalconfig.ViewFuncConfig,
			appconfig.TemplateAutoLoad,
		),
	}

	for _, cbFunc := range app.conf.DiBindings() {
		key, binding, err := cbFunc(app.di)
		if err != nil {
			panic(err.Error())
		}
		app.di.Set(key, binding)
	}
	app.di.Set("internal.app.config.Configer", app.conf)
	app.di.Set("olbrichattila.godi.Container", app.di)

	_, err := app.di.Get(app)
	if err != nil {
		panic(err.Error())
	}

	_, err = app.di.Call(Bootstrap.Bootstrap)
	if err != nil {
		panic(err.Error())
	}

	return app
}

type App struct {
	di              godi.Container
	router          router.Router
	conf            config.Configer
	commandExecutor commandexecutor.CommandExecutor
}

func (a *App) Construct(
	_ env.Enver, // It will automatically loads env with it's constructor
	cron cron.JobTimer,
	router router.Router,
	ce commandexecutor.CommandExecutor,
) {
	cron.Init(a.di, a.conf.Jobs())
	a.router = router
	a.commandExecutor = ce
}

func (a *App) Serve() {
	port, err := a.getPort()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	hTTPHandler := &hTTPHandler{app: a}
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.Handle("/", hTTPHandler)

	// TODO Add this in go routine to listen on https as well
	// http.ListenAndServeTLS()
	err = http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (a *App) getPort() (string, error) {
	port := os.Getenv("HTTP_LISTENING_PORT")
	if port == "" {
		return ":80", nil
	}

	if _, err := strconv.Atoi(port); err == nil {
		return ":" + port, nil
	}

	return "", fmt.Errorf("port %s provided is not a number", port)
}

func (a *App) Command() {
	err := a.commandExecutor.Execute(a.di, a.conf.ConsoleCommands())
	if err != nil {
		fmt.Println(err.Error())
	}
}

type hTTPHandler struct {
	app *App
}

func (h *hTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	routes := h.app.conf.Routes()
	customValidator := h.getValidatorFromDi()
	if customValidator != nil {
		customValidator.SetRules(appconfig.ValidatorRules)
		customValidator.SetRules(internalconfig.ValidatorRules)
	}
	session := h.getSessionerFromDi()
	req := h.getRequestFromDi()
	if req != nil {
		req.SetRequest(r)
	}

	h.app.di.Set("http.ResponseWriter", w)
	for _, middleware := range h.app.conf.Middlewares() {
		res, err := h.app.di.Call(middleware)
		if err != nil {
			panic(err.Error())
		}
		if len(res) > 0 && res[0].Kind() == reflect.Bool {
			if !res[0].Bool() {
				return
			}
		}
	}

	for _, action := range routes {
		match, routePars := h.app.router.Match(action.Path, r.RequestURI)

		if match {
			if action.RequestType != r.Method {
				continue
			}

			if req != nil {
				req.SetRouteParameters(routePars)
			}

			for _, rootMiddleware := range action.Middlewares {
				res, err := h.app.di.Call(rootMiddleware)
				if err != nil {
					panic(err.Error())
				}
				if len(res) > 0 && res[0].Kind() == reflect.Bool {
					if !res[0].Bool() {
						return
					}
				}
			}

			// Route validator logic
			if action.ValidationRules != "" {
				genericErrors := make(validator.ValidationErrors)
				funcErrors := make(validator.ValidationErrors)
				isValid := true
				if rule, ok := appconfig.RouteValidationRules[action.ValidationRules]; ok {

					if customValidator != nil {
						allRequests := req.AllFlat()
						if rule.Rules != nil {

							ok, errors, _ := customValidator.Validate(allRequests, rule.Rules)
							if !ok {
								genericErrors = errors
								isValid = false
							}
						}

						if rule.CustomRule != nil {
							if customFuncErrors, ok := rule.CustomRule(allRequests); !ok {
								funcErrors = customFuncErrors
								isValid = false
							}
						}

						if !isValid {
							if session != nil {
								combinedErrors := h.mergeValidationErrors(genericErrors, funcErrors)
								jSONError, err := json.Marshal(combinedErrors)
								if err == nil {
									session.Set("lastValidationError", string(jSONError))
								}
							}

							if rule.Redirect != "" {
								http.Redirect(w, r, rule.Redirect, http.StatusSeeOther)
								return
							}
						}
					}
				}
			}

			// This is the main controller call
			result, err := h.app.di.Call(action.Fn)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}

			if len(result) == 2 {
				errorInterface := reflect.TypeOf((*error)(nil)).Elem()
				if result[1].Type().Implements(errorInterface) {
					// Use type assertion to get the error
					if err, ok := result[1].Interface().(error); ok {
						w.WriteHeader(http.StatusInternalServerError)
						w.Write([]byte(err.Error()))
						return
					}
				}
			}

			if len(result) == 0 {
				return
			}

			if result[0].Kind() == reflect.String {
				w.Write([]byte(result[0].String()))
				return
			}

			if result[0].Kind() == reflect.Struct || result[0].Kind() == reflect.Map {
				jsonRes, err := json.Marshal(result[0].Interface())
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte(err.Error()))
					return
				}
				w.Header().Set("Content-Type", "application/json")
				w.Write(jsonRes)
				return
			}

		}
	}

	http.NotFound(w, r)
}

func (h *hTTPHandler) getRequestFromDi() request.Requester {
	dep, err := h.app.di.GetDependency("internal.app.request.Requester")
	if err == nil {
		if req, ok := dep.(request.Requester); ok {
			return req
		}
	}

	return nil
}

func (h *hTTPHandler) getValidatorFromDi() validator.Validator {
	dep, err := h.app.di.GetDependency("internal.app.validator.Validator")
	if err == nil {
		if req, ok := dep.(validator.Validator); ok {
			return req
		}
	}

	return nil
}

func (h *hTTPHandler) getSessionerFromDi() session.Sessioner {
	dep, err := h.app.di.GetDependency("internal.app.session.Sessioner")
	if err == nil {
		if req, ok := dep.(session.Sessioner); ok {
			return req
		}
	}

	return nil
}

func (h *hTTPHandler) mergeValidationErrors(errorSet1, errorSet2 validator.ValidationErrors) validator.ValidationErrors {
	result := make(validator.ValidationErrors)
	for key, value := range errorSet1 {
		if value == nil {
			result[key] = make([]string, 0)
			continue
		}

		result[key] = value
	}

	for key, value := range errorSet2 {
		subset, ok := result[key]
		if ok && value != nil {
			result[key] = append(subset, value...)
			continue
		}

		if value == nil {
			result[key] = make([]string, 0)
			continue
		}

		result[key] = value

	}
	return result
}
