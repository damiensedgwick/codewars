package main

import "strings"

func IncrementOne(n int) int {
	return n + 1
}

func DecrementOne(n int) int {
	return n - 1
}

func Square(n int) int {
	return n * n
}

func SplitString(s string) []string {
	return strings.Split(s, "")
}

func Parse(input string) []int {
	var output = []int{}

	count := 0

	s := SplitString(input)
	for _, v := range s {
		switch v {
		case "i":
			count = IncrementOne(count)
		case "d":
			count = DecrementOne(count)
		case "s":
			count = Square(count)
		case "o":
			output = append(output, count)
		default:
			continue
		}
	}

	return output
}
