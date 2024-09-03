package internalconfig

import (
	internalviewfunction "framework/internal/app/view-functions"
	"html/template"
	"net/url"
	"os"
)

var ViewFuncConfig = template.FuncMap{
	"urlEscape":    url.QueryEscape,
	"envVar":       os.Getenv,
	"renderErrors": internalviewfunction.RenderErrors,
	"renderError":  internalviewfunction.RenderError,
	"lastRequest":  internalviewfunction.RenderRequest,
}
