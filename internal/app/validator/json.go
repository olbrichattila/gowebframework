package validator

import (
	"encoding/json"
)

func jsonRule(val string, _ ...string) (string, bool) {
	var js json.RawMessage
	if json.Unmarshal([]byte(val), &js) == nil {
		return "", true
	}

	return "not valid JSON", false
}
