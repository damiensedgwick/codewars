// link to kata https://www.codewars.com/kata/525f50e3b73515a6db000b83/train/go
package main

import (
	"fmt"
	"strconv"
)

func MakeStringSlice(intSlice []int) []string {
	var slice []string
	for _, i := range intSlice {
		slice = append(slice, strconv.Itoa(i))
	}

	return slice
}

func main() {
	fmt.Println("Creating phone numbers")
}
