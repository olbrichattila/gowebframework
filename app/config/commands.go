package appconfig

import (
	command "framework/app/commands"
	commandexecutor "framework/internal/app/command"
)

var ConsoleCommands = map[string]commandexecutor.CommandItem{
	"command1": {Fn: command.Test, Desc: ""},
}
