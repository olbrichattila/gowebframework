package validator

type ruleFunc func(string, ...string) (string, bool)

var ruleMap = map[string]ruleFunc{
	"required": requiredRule,
	"min":      minRule,
	"max":      maxRule,
	"in":       inRule,
	"regex":    regexRule,
	"between":  betweenRule,
	"size":     sizeRule,
	"email":    emailRule,
	"url":      urlRule,
	"uuid":     uuidRule,
	"numeric":  numericRule,
	"integer":  integerRule,
	"date":     dateRule,
	"dateTime": dateTimeRule,
	"boolean":  booleanRule,
	"json":     jsonRule,
}
