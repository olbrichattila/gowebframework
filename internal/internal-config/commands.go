package internalconfig

import (
	commandexecutor "framework/internal/app/command"
	internalcommand "framework/internal/internal-commands"
)

var ConsoleCommands = map[string]commandexecutor.CommandItem{
	"list-commands":         {Fn: internalcommand.ListCommands, Desc: ""},
	"create:command":        {Fn: internalcommand.CreateCommand, Desc: ""},
	"create:controller":     {Fn: internalcommand.CreateController, Desc: "possible flags: (-api, -rest -in= -out=). try -help for more details"},
	"create:middleware":     {Fn: internalcommand.CreateMiddleware, Desc: "possible flags: (-in= -out=). try -help for more details"},
	"create:job":            {Fn: internalcommand.CreateJob, Desc: "possible flags: (-in= -out=). try -help for more details"},
	"create:event-consumer": {Fn: internalcommand.CreateEventConsumer, Desc: ""},
}
