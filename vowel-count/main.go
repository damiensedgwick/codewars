// link to kata https://www.codewars.com/kata/54ff3102c1bad923760001f3/train/go
package main

import (
	"fmt"
	"strings"
)

func GetCount(str string) (count int) {
	slice := strings.Split(str, "")
	vowels := []string{"a", "e", "i", "o", "u"}

	count = 0
	for i := 0; i < len(slice); i++ {
		if Contains(vowels, slice[i]) {
			count++
		}
	}

	return count
}

func Contains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}

	return false
}

func main() {
	c := GetCount("abracadabra")

	fmt.Println(c)
}
