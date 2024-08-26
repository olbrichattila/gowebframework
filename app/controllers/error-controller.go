package controller

import (
	"fmt"
	"framework/internal/app/session"
	"framework/internal/app/validator"
	"framework/internal/app/view"
)

func DisplayError(v view.Viewer, s session.Sessioner, val validator.Validator) string {
	values := map[string]string{
		"fieldName":  "65",
		"fieldName2": "54",
		"fieldName3": "hello",
	}
	rules := map[string]string{
		"fieldName":  "min:55",
		"fieldName2": "max:55",
		"fieldName3": "in:a,bc,de,hello,bukk",
	}
	ok, messages, validated := val.Validate(values, rules)
	fmt.Printf("%v-%v-%v", ok, messages, validated)
	data := map[string]string{
		"lastError": s.Get("lastError"),
	}

	return v.RenderView("error.html", data)
}
