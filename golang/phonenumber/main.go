// link to kata https://www.codewars.com/kata/525f50e3b73515a6db000b83/train/go
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func MakeStringSlice(digits []int) []string {
	var s []string
	for _, i := range digits {
		s = append(s, strconv.Itoa(i))
	}

	return s
}

func CreateAreaCode(s []string) string {
	ac := "(" + strings.Join(s[:3], "") + ")"

	return ac
}

func CreatePhoneNumberPrefix(s []string) string {
	p := strings.Join(s[3:6], "") + "-"

	return p
}

func ConcatRemainingDigits(s []string) string {
	rd := strings.Join(s[6:], "")

	return rd
}

func PhoneNumberBuilder(code, prefix, remaining string) string {
	pn := code + " " + prefix + remaining

	return pn
}

// CreatePhoneNumber([10]uint{1,2,3,4,5,6,7,8,9,0})
// Should return "(123) 456-7890"

func main() {
	fmt.Println("Creating phone numbers")
}
