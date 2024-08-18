package command

import (
	"fmt"
	"framework/internal/app/args"
	"strings"
)

func Test(a args.CommandArger) {
	fmt.Println(strings.Join(a.GetAll(), ","))

	arg, err := a.Get(1)
	fmt.Println(arg, err)

	fmt.Println(a.GetAllFlags())

	flagValue, ok := a.GetFlagByName("hello", "default")

	fmt.Println(flagValue, ok)

	flagValue, ok = a.GetFlagByName("nonexistent", "default")

	fmt.Println(flagValue, ok)
}
