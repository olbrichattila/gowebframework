package controller

import (
	"framework/internal/app/session"
	"framework/internal/app/validator"
	"framework/internal/app/view"
)

func DisplayError(v view.Viewer, s session.Sessioner, val validator.Validator) string {
	data := map[string]string{
		"lastError": s.Get("lastError"),
	}

	return v.RenderView("error.html", data)
}
