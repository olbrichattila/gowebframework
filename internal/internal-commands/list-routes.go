package internalcommand

import (
	"fmt"
	"framework/internal/app/config"
	"framework/internal/app/router"
	"sort"
)

func ListRoutes(c config.Configer) {
	routes := c.Routes()
	reorderSlice := make([]router.ControllerAction, len(routes))
	copy(reorderSlice, routes)

	sort.Slice(reorderSlice, func(i, j int) bool {
		return reorderSlice[i].Path < reorderSlice[j].Path
	})

	for _, route := range reorderSlice {
		fmt.Println(route.Path)
	}
}
