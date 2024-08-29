package internalcommand

import (
	"fmt"
	"framework/internal/app/config"
	"reflect"
	"runtime"
)

func ListMiddlewares(c config.Configer) {
	middlewares := c.Middlewares()
	reorderSlice := make([]string, len(middlewares))

	for i, middleware := range middlewares {
		reorderSlice[i] = runtime.FuncForPC(reflect.ValueOf(middleware).Pointer()).Name()
	}

	for _, middlewareName := range reorderSlice {
		fmt.Println(middlewareName)
	}
}
