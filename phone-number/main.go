// link to kata https://www.codewars.com/kata/525f50e3b73515a6db000b83/train/go
package main

import (
	"strconv"
	"strings"
)

func MakeStringSlice(digits [10]uint) []string {
	var s []string
	for _, i := range digits {
		s = append(s, strconv.FormatUint(uint64(i), 10))
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

func CreatePhoneNumber(numbers [10]uint) string {
	strSlice := MakeStringSlice(numbers)
	code := CreateAreaCode(strSlice)
	prefix := CreatePhoneNumberPrefix(strSlice)
	remaining := ConcatRemainingDigits(strSlice)
	phoneNumber := PhoneNumberBuilder(code, prefix, remaining)

	return phoneNumber
}
