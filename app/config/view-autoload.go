package appconfig

import "framework/internal/app/view"

var TemplateAutoLoad = map[string][]string{
	view.ViewTypeHTML: {
		"template/head.html",
		"template/header.html",
		"template/footer.html",
	},
	view.ViewTypeEmail: {},
}
