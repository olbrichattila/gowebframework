package internalviewfunction

import (
	"encoding/json"
	"framework/internal/app/validator"
	"strings"
)

func RenderErrors(jsonStr string) string {
	var validatorErrors validator.ValidationErrors
	err := json.Unmarshal([]byte(jsonStr), &validatorErrors)
	if err != nil {
		return err.Error()
	}

	sw := &strings.Builder{}
	sw.WriteString("<ul>")
	for field, errorList := range validatorErrors {
		sw.WriteString("<li>")
		sw.WriteString(field)
		sw.WriteString("<ul>")
		for _, eMess := range errorList {
			sw.WriteString("<li>")
			sw.WriteString(eMess)
			sw.WriteString("</li>")
		}
		sw.WriteString("</ul>")
		sw.WriteString("</li>")
	}
	sw.WriteString("</ul>")

	return sw.String()
}

func RenderError(field, jsonStr string) string {
	var validatorErrors validator.ValidationErrors

	err := json.Unmarshal([]byte(jsonStr), &validatorErrors)
	if err != nil {
		return err.Error()
	}

	if errorList, ok := validatorErrors[field]; ok {
		return strings.Join(errorList, ", ")
	}

	return ""
}
