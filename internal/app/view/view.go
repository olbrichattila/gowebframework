package view

import (
	"bytes"
	"framework/internal/app/config"
	"os"
	"path/filepath"
	"text/template"
)

type Viewer interface {
	Construct(config.Configer)
	Render([]string, any) string
	RenderMail([]string, any) string
	NewPath(...string)
	RenderToFile(string, []string, any) error
	RenderMailToFile(string, []string, any) error
	Funcs(template.FuncMap) Viewer
}

func New() Viewer {
	return &View{
		path: []string{"app", "views"},
	}
}

type View struct {
	config config.Configer
	path   []string
	funcs  template.FuncMap
}

func (v *View) Construct(conf config.Configer) {
	v.config = conf
}

func (v *View) RenderToFile(fileName string, templates []string, params any) error {
	content := v.Render(templates, params)
	return v.toFile(fileName, content)
}

func (v *View) RenderMailToFile(fileName string, templates []string, params any) error {
	content := v.RenderMail(templates, params)
	return v.toFile(fileName, content)
}

func (v *View) toFile(fileName, content string) error {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}

func (v *View) Render(templates []string, params any) string {
	if len(templates) == 0 {
		return ""
	}

	paths := v.addPath(templates)
	funcs := v.mergeFuncMap()
	tmpl, err := template.New("example").Funcs(funcs).ParseFiles(paths...)
	if err != nil {
		return err.Error()
	}

	// Create a bytes.Buffer to hold the rendered template
	var buf bytes.Buffer
	err = tmpl.ExecuteTemplate(&buf, templates[0], params)
	if err != nil {
		return err.Error()
	}
	v.funcs = make(template.FuncMap, 0)

	// Return the rendered template as a string
	return buf.String()
}

func (v *View) RenderMail(templates []string, params any) string {
	v.NewPath("app", "mails")
	result := v.Render(templates, params)
	v.NewPath("app", "views")
	return result
}

func (v *View) Funcs(fm template.FuncMap) Viewer {
	v.funcs = fm
	return v
}

func (v *View) addPath(templates []string) []string {
	result := make([]string, len(templates))
	viewsDir := filepath.Join(v.path...)

	for i, templateFileName := range templates {
		result[i] = viewsDir + "/" + templateFileName
	}

	return result
}

func (v *View) NewPath(p ...string) {
	v.path = p
}

func (v *View) mergeFuncMap() template.FuncMap {
	merged := make(template.FuncMap, 0)
	for funcName, value := range v.config.ViewConfig() {
		merged[funcName] = value
	}

	for funcName, value := range v.funcs {
		merged[funcName] = value
	}

	return merged
}
