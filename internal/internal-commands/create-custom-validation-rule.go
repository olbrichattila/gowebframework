package internalcommand

import (
	"fmt"
	"framework/internal/app/args"
	commandcreator "framework/internal/app/wizards/command"
)

func CreateCustomValidationRule(a args.CommandArger, c commandcreator.CommandCreator) {
	err := c.Create("customrule.tpl", "./app/validator-configs", nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Please register your new rule in in:\n app/config/validators.go\n")
}
