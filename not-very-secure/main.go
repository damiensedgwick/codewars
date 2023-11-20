// lint to kata https://www.codewars.com/kata/526dbd6c8c0eb53254000110/train/go
package main

import "regexp"

func Alphanumeric(str string) bool {
	return regexp.MustCompile("^[a-zA-Z0-9]+$").MatchString(str)
}
