// link to kata https://www.codewars.com/kata/541c8630095125aba6000c00/train/go
package main

import (
	"strconv"
	"strings"
)

func MakeSlice(n int) []int {
	var s []int
	str := strconv.Itoa(n)
	strSlice := strings.Split(str, "")
	for _, v := range strSlice {
		num, _ := strconv.Atoi(v)
		s = append(s, num)
	}

	return s
}

func SumSliceDigits(s []int) int {
	sum := 0
	for _, v := range s {
		sum += v
	}

	if len(MakeSlice(sum)) == 1 {
		return sum
	}

	return SumSliceDigits(MakeSlice(sum))
}

func DigitalRoot(n int) int {
	return SumSliceDigits(MakeSlice(n))
}
