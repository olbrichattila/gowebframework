package internalcommand

import (
	"fmt"
	"framework/internal/app/args"
	commandcreator "framework/internal/app/wizards/command"
)

func CreateCommand(a args.CommandArger, c commandcreator.CommandCreator) {
	err := c.Create("command.tpl", "./app/commands", nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Please register your new command in:\n  app/config/commands.go\n")
}
