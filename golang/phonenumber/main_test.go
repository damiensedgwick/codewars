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
	}, MakeStringSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
}
