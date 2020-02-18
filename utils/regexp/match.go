package regexp

import "regexp"

var phone *regexp.Regexp
var code *regexp.Regexp
var stuid *regexp.Regexp
var username *regexp.Regexp

// bound for consideration about speed by golang, there is not regexp which can done this work at once.
// Go's regex doesn't support backtracking.
// https://stackoverflow.com/questions/25837241/password-validation-with-regexp
var passwordOne *regexp.Regexp
var passwordTwo *regexp.Regexp
var token *regexp.Regexp

func init() {
	phone = regexp.MustCompile(`^1\d{10}$`)
	code = regexp.MustCompile(`^\d{4}$`)
	stuid = regexp.MustCompile(`^2017\d{8}(x)?$`)
	username = regexp.MustCompile(`^[[:alnum:]]{2,16}$`)
	passwordOne = regexp.MustCompile(`^[[:alnum:]]{8,16}$`)
	passwordTwo = regexp.MustCompile(`[A-Z]+|[a-z]+`)
	token = regexp.MustCompile(`^[A-Za-z0-9+/]+$`)
}

func RegexpPhone(str string) bool {
	return phone.MatchString(str)
}

func RegexpCode(str string) bool {
	return code.MatchString(str)
}

func RegexpStuid(str string) bool {
	return stuid.MatchString(str)
}

func RegexpUsername(str string) bool {
	return username.MatchString(str)
}

func RegexpPassword(str string) bool {
	return passwordOne.MatchString(str) && passwordTwo.MatchString(str)
}

func RegexpToken(str string) bool {
	return token.MatchString(str)
}
