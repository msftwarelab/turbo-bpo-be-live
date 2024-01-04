package utils

import (
	"regexp"
)

// Test if a string is a valid email
// param str - A string to test
// return True if correct, false otherwise
func IsEmail(str string) bool {
	exp := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return exp.MatchString(str) == true
}

func Positive(val int64) int64 {
	if val < 0 {
		return -1 * val
	} else {
		return val
	}
}
