package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMakeStringSlice(t *testing.T) {
	assert.Equal(t, []string{
		"1",
		"2",
		"3",
		"4",
		"5",
		"6",
		"7",
		"8",
		"9",
		"0",
	}, MakeStringSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}))
}

func TestCreateAreaCode(t *testing.T) {
	assert.Equal(t, "(123)", CreateAreaCode([]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}))
}

func TestCreatePhoneNumberPrefix(t *testing.T) {
	assert.Equal(t, "456-", CreatePhoneNumberPrefix([]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}))
}

func TestConcatRemainingDigits(t *testing.T) {
	assert.Equal(t, "7890", ConcatRemainingDigits([]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}))
}

func TestPhoneNumberBuilder(t *testing.T) {
	assert.Equal(t, "(123) 456-7890", PhoneNumberBuilder("(123)", "456-", "7890"))
}
