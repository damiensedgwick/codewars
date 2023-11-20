// link to kata https://www.codewars.com/kata/587731fda577b3d1b0001196
package main

import (
	"strings"
)

func MakeStringSlice(s string) []string {
	return strings.Split(s, " ")
}

func CapitaliseEachStringInSlice(s []string) []string {
	var capitalisedSlice []string

	for _, v := range s {
		capitalisedSlice = append(capitalisedSlice, strings.Title(v))
	}

	return capitalisedSlice
}

func MakeCamelCaseString(s []string) string {
	return strings.Join(s, "")
}

func CamelCase(s string) string {
	strSlice := MakeStringSlice(s)
	capitalisedStrSlice := CapitaliseEachStringInSlice(strSlice)
	camelCaseStr := MakeCamelCaseString(capitalisedStrSlice)

	return camelCaseStr
}
