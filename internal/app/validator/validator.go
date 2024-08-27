package validator

import (
	"strings"
)

const (
	pipePlaceholder = "__PIPE__PLACEHOLDER__"
)

func New() Validator {
	return &Validate{}
}

type Validator interface {
	Validate(map[string]string, map[string]string) (bool, []string, map[string]string)
}

type Validate struct {
}

func (v *Validate) Validate(fields map[string]string, rules map[string]string) (bool, []string, map[string]string) {
	valid := make(map[string]string)
	validationErrors := make([]string, 0)
	for field, rule := range rules {
		fieldValue := fields[field]
		if message, ok := v.parse(field, fieldValue, rule); ok {
			valid[field] = fieldValue
		} else {
			validationErrors = append(validationErrors, message...)
		}
	}

	return len(validationErrors) == 0, validationErrors, valid
}

func (v *Validate) parse(fieldName, val, pattern string) ([]string, bool) {
	errorMessages := make([]string, 0)
	pattern = strings.ReplaceAll(pattern, `\|`, pipePlaceholder)
	patterns := strings.Split(pattern, "|")
	for _, rule := range patterns {
		rule = strings.ReplaceAll(rule, pipePlaceholder, "|")
		if message, ok := v.parseRule(val, rule); !ok {
			if message != "" {
				message = fieldName + ": " + message
			}
			errorMessages = append(errorMessages, message)
		}
	}
	return errorMessages, len(errorMessages) == 0
}

func (*Validate) parseRule(val, rule string) (string, bool) {
	rulePars := make([]string, 0)
	ruleParts := strings.Split(rule, ":")
	ruleName := ruleParts[0]
	if len(ruleParts) > 1 {
		rulePars = strings.Split(ruleParts[1], ",")
	}

	if ruleFn, ok := ruleMap[ruleName]; ok {
		return ruleFn(val, rulePars...)
	}

	return "", true
}
