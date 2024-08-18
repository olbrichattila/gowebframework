package commandcreator

import (
	"fmt"
	"framework/internal/app/args"
	"framework/internal/app/view"
	"os"
	"regexp"
	"strings"
)

func New() CommandCreator {
	return &Creator{}
}

type CommandCreator interface {
	Construct(a args.CommandArger, v view.Viewer)
	Create(string, string, map[string]string) error
}

type Creator struct {
	a args.CommandArger
	v view.Viewer
}

func (c *Creator) Construct(a args.CommandArger, v view.Viewer) {
	c.a = a
	c.v = v
}

func (c *Creator) Create(templateName, savePath string, data map[string]string) error {
	templateFiles := []string{templateName}
	commandName, err := c.a.Get(0)
	if err != nil {
		return fmt.Errorf("file name not provided")
	}

	fileName := fmt.Sprintf("%s/%s.go", savePath, commandName)
	if _, err := os.Stat(fileName); err == nil {
		return fmt.Errorf("file already exists")
	}

	mergedData := map[string]string{
		"name": c.cleanFileName(commandName),
	}

	for key, value := range data {
		mergedData[key] = value
	}

	c.v.NewPath("internal", "command-templates")
	err = c.v.RenderToFile(fileName, templateFiles, mergedData)
	if err != nil {
		return err
	}

	return nil
}

func (c *Creator) cleanFileName(fn string) string {
	sb := &strings.Builder{}
	words := strings.FieldsFunc(fn, func(r rune) bool {
		return r == ' ' || r == '-' || r == '_'
	})

	for _, word := range words {
		sb.WriteString(
			c.filterSpecialChars(
				c.lcFirst(word),
			),
		)
	}

	return sb.String()

}

func (*Creator) lcFirst(s string) string {
	return strings.ToUpper(string(s[0])) + s[1:]
}

func (*Creator) filterSpecialChars(s string) string {

	re := regexp.MustCompile("[^a-zA-Z0-9]+")

	return re.ReplaceAllString(s, "")

}
