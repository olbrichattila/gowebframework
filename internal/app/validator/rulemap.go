package validator

type ruleFunc func(string, ...string) (string, bool)

var ruleMap = map[string]ruleFunc{
	"required": requiredRule,
	"min":      minRule,
	"max":      maxRule,
	"in":       inRule,
	"regex":    regexRule,
	"between":  betweenRule,
}
