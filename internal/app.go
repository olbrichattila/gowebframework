// Package app
package app

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
	internalconfig "framework/internal/internal-config"
	"net/http"
	"reflect"

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
	conf            config.Configer
	commandExecutor commandexecutor.CommandExecutor
}

func (a *App) Construct(
	_ env.Enver, // It will automatically loads env with it's constructor
	cron cron.JobTimer,
	ce commandexecutor.CommandExecutor,
) {
	cron.Init(a.di, a.conf.Jobs())
	a.commandExecutor = ce
}

func (a *App) Serve() {
	hTTPHandler := &hTTPHandler{app: a}
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.Handle("/", hTTPHandler)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
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
	dep, err := h.app.di.GetDependency("internal.app.request.Requester")
	if err == nil {
		if req, ok := dep.(request.Requester); ok {
			req.SetRequest(r)
		}

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
		if action.Path == r.URL.Path {
			if action.RequestType != r.Method {
				continue
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
