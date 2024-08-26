package appconfig

import "framework/internal/app/view"

var TemplateAutoLoad = map[string][]string{
	// Template partials rendered by v.RenderView, template folder is under views
	view.ViewTypeHTML: {
		"template/head.html",
		"template/header.html",
		"template/footer.html",
	},
	// Template partials rendered by v.RenderMail, template folder is under mails
	view.ViewTypeEmail: {
		"template/header.html",
		"template/footer.html",
	},
}
