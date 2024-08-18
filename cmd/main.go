package main

import (
	"fmt"
	app "framework/internal"
	"os"

	"github.com/olbrichattila/godi"
)

func main() {
	args := os.Args

	container := godi.New()
	app := app.New(container)
	if len(args) < 2 {
		app.Serve()
		return
	}

	switch args[1] {
	case "serve":
		app.Serve()
	case "artisan":
		app.Command()
	default:
		displayHelp()
		return
	}
}

func displayHelp() {
	fmt.Printf(
		`Usage:
Run HTTP server:	
     go run ./cmd serve
Run command:	 
     go run ./cmd artisan <command>
`,
	)
}
