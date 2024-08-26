package validator

import (
	"strings"
)

func requiredRule(val string, _ ...string) (string, bool) {
	if len(strings.TrimSpace(val)) == 0 {
		return "required", false
	}

	return "", true

}
