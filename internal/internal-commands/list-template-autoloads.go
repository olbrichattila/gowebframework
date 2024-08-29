package internalcommand

import (
	"fmt"
	"framework/internal/app/config"
	"sort"
)

func ListTemplateAutoLoads(c config.Configer) {
	templateAutoLoads := c.GetTemplateAutoLoads()

	keys := make([]string, 0, len(templateAutoLoads))
	for k := range templateAutoLoads {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	for _, name := range keys {
		items := templateAutoLoads[name]
		fmt.Printf("- %s\n", name)
		for _, item := range items {
			fmt.Printf("   %s\n", item)
		}
	}
}
