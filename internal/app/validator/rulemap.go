package validator

type ruleFunc func(string, ...string) (string, bool)

var ruleMap = map[string]ruleFunc{
	"min": minRule,
	"max": maxRule,
	"in":  inRule,
}
