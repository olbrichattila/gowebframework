package validator

import (
	"encoding/json"
	"fmt"
)

func jsonRule(val string, _ ...string) (string, bool) {
	var js json.RawMessage
	if json.Unmarshal([]byte(val), &js) == nil {
		return "", true
	}

	return fmt.Sprintf("%s not in an JSON", val), false
}
