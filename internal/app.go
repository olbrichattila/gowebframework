// Package app
package app

// TODO: Refactor, split up, getting too big
import (
	"fmt"
	Bootstrap "framework/app"
	appconfig "framework/app/config"
	commandexecutor "framework/internal/app/command"
	"framework/internal/app/config"
	"framework/internal/app/cron"
	"framework/internal/app/env"
	"framework/internal/app/router"
	internalconfig "framework/internal/internal-config"
	"net/http"
	"os"
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

	app.initBindings()

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

func (a *App) initBindings() {
	for _, cbFunc := range a.conf.DiBindings() {
		key, binding, err := cbFunc(a.di)
		if err != nil {
			panic(err.Error())
		}
		a.di.Set(key, binding)
	}
	a.di.Set("internal.app.config.Configer", a.conf)
	a.di.Set("olbrichattila.godi.Container", a.di)
}
