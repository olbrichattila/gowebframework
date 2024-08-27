package appconfig

type CustomValidatorFunc func(map[string]string) (string, bool)

type ValidationRule struct {
	Redirect   string
	Rules      map[string]string
	CustomRule CustomValidatorFunc
}

var RouteValidationRules = map[string]ValidationRule{
	"register": {
		Redirect: "/register",
		Rules: map[string]string{
			"password": "minSize:6|maxSize:255",
			"name":     "minSize:6|maxSize:255",
			"email":    "email",
		},
		// CustomRule: func(fields map[string]string) (string, bool) { return "not good " + fmt.Sprintf("%v", fields), false },
	},
}
