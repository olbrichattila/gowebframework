package internalconfig

import (
	"html/template"
	"net/url"
)

var ViewFuncConfig = template.FuncMap{
	"urlEscape": url.QueryEscape,
}
