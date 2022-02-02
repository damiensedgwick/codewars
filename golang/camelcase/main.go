package main

import (
	"fmt"
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

func main() {
	fmt.Println(CamelCase("camel case string"))
}
