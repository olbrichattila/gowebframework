package config

import (
	commandexecutor "framework/internal/app/command"
	"framework/internal/app/cron"
	"framework/internal/app/router"

	"github.com/olbrichattila/godi"
)

type DiCallback func(godi.Container) (string, interface{}, error)

func New(
	routes []router.ControllerAction,
	jobs []cron.Job,
	middlewares []interface{},
	appBindings []DiCallback,
	internalBindings []DiCallback,
	appCommands map[string]commandexecutor.CommandItem,
	internalCommands map[string]commandexecutor.CommandItem,
) Configer {
	return &Conf{
		routes:           routes,
		jobs:             jobs,
		middlewares:      middlewares,
		appBindings:      appBindings,
		internalBindings: internalBindings,
		appCommands:      appCommands,
		internalCommands: internalCommands,
	}
}

type Configer interface {
	Routes() []router.ControllerAction
	DiBindings() []DiCallback
	ConsoleCommands() map[string]commandexecutor.CommandItem
	Jobs() []cron.Job
	Middlewares() []interface{}
}

type Conf struct {
	routes           []router.ControllerAction
	jobs             []cron.Job
	middlewares      []interface{}
	appBindings      []DiCallback
	internalBindings []DiCallback
	appCommands      map[string]commandexecutor.CommandItem
	internalCommands map[string]commandexecutor.CommandItem
}

func (c *Conf) Routes() []router.ControllerAction {
	return c.routes
}

func (c *Conf) DiBindings() []DiCallback {
	return append(c.appBindings, c.internalBindings...)
}

func (c *Conf) ConsoleCommands() map[string]commandexecutor.CommandItem {
	return c.mergeCommands(c.appCommands, c.internalCommands)
}

func (c *Conf) Jobs() []cron.Job {
	return c.jobs
}

func (c *Conf) Middlewares() []interface{} {
	return c.middlewares
}

func (*Conf) mergeCommands(maps ...map[string]commandexecutor.CommandItem) map[string]commandexecutor.CommandItem {
	mergedMap := make(map[string]commandexecutor.CommandItem)

	for _, mergeMap := range maps {
		for key, value := range mergeMap {
			mergedMap[key] = value
		}
	}

	return mergedMap
}
