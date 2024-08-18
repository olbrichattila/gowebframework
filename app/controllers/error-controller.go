package controller

import (
	"framework/internal/app/session"
	"framework/internal/app/view"
)

func DisplayError(v view.Viewer, s session.Sessioner) string {
	templateFiles := []string{
		"error.html",
	}

	data := map[string]string{
		"lastError": s.Get("lastError"),
	}

	return v.Render(templateFiles, data)
}
