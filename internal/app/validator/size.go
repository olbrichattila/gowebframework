package validator

import (
	"fmt"
	"strconv"
)

func sizeRule(val string, pars ...string) (string, bool) {
	if len(pars) != 1 {
		return "requires 1 numeric parameter like size:1", false
	}

	a, err := strconv.Atoi(val)
	if err != nil {
		return "The value comparing must be a number", false
	}

	if len(val) == a {
		return "", true
	}

	return fmt.Sprintf("%s is not %d long", val, a), false
}
