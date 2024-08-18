package view

import (
	"bytes"
	"os"
	"path/filepath"
	"text/template"
)

type Viewer interface {
	Render([]string, any) string
	RenderMail([]string, any) string
	NewPath(...string)
	RenderToFile(string, []string, any) error
	RenderMailToFile(string, []string, any) error
}

func New() Viewer {
	return &View{
		path: []string{"app", "views"},
	}
}

type View struct {
	path []string
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
	tmpl, err := template.ParseFiles(paths...)
	if err != nil {
		return err.Error()
	}

	// Create a bytes.Buffer to hold the rendered template
	var buf bytes.Buffer

	err = tmpl.ExecuteTemplate(&buf, templates[0], params)
	if err != nil {
		return err.Error()
	}

	// Return the rendered template as a string
	return buf.String()
}

func (v *View) RenderMail(templates []string, params any) string {
	v.NewPath("app", "mails")
	result := v.Render(templates, params)
	v.NewPath("app", "views")
	return result
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
