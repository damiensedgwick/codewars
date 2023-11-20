package main

import (
	"strconv"
	"strings"
)

func ValidateInput(input string) bool {
	octets := strings.Split(input, ".")
	if len(octets) != 4 {
		return false
	}

	for _, v := range octets {
		s := strings.Split(v, "")
		if s[0] == "0" {
			return false
		}

		num, err := strconv.Atoi(v)
		if err != nil {
			return false
		}

		if num > 255 {
			return false
		}

	}

	return true
}

func Is_valid_ip(ip string) bool {
	return ValidateInput(ip)
}

// Best codewars solution
// func Is_valid_ip(ip string) bool {
// 	return net.ParseIP(ip) != nil
// }
