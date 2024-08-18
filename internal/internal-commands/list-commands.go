package internalcommand

import (
	"fmt"
	"framework/internal/app/config"
	"sort"
)

func ListCommands(c config.Configer) {
	commands := c.ConsoleCommands()

	keys := make([]string, 0, len(commands))
	for k := range commands {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, name := range keys {
		commandItem := commands[name]
		fmt.Printf("- %s\n", name)
		if commandItem.Desc != "" {
			fmt.Printf("     %s\n", commandItem.Desc)
		}
	}

}
