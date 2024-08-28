package internalconfig

import (
	internalviewfunction "framework/internal/app/view-functions"
	"html/template"
	"net/url"
)

var ViewFuncConfig = template.FuncMap{
	"urlEscape":    url.QueryEscape,
	"renderErrors": internalviewfunction.RenderErrors,
	"renderError":  internalviewfunction.RenderError,
}
