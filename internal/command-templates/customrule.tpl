package customrule

// {{.name}}Rule is a custom validator rule, 
// val is the value to validate, 
// pars is the elements in the rule signature, like myrule:1,2,3 will be 1, 2 and 3
// returns error message and bool if validation is OK
func {{.name}}Rule(val string, pars ...string) (string, bool) {
    return "", true
}
