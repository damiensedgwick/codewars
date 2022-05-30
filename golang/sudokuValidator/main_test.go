package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var puzzle = [][]int{
	{5, 3, 4, 6, 7, 8, 9, 1, 2},
	{6, 7, 2, 1, 9, 5, 3, 4, 8},
	{1, 9, 8, 3, 4, 2, 5, 6, 7},
	{8, 5, 9, 7, 6, 1, 4, 2, 3},
	{4, 2, 6, 8, 5, 3, 7, 9, 1},
	{7, 1, 3, 9, 2, 4, 8, 5, 6},
	{9, 6, 1, 5, 3, 7, 2, 8, 4},
	{2, 8, 7, 4, 1, 9, 6, 3, 5},
	{3, 4, 5, 2, 8, 6, 1, 7, 9},
}

// grab the first row
func TestCheckForZero(t *testing.T) {
	assert.Equal(t, false, CheckForZero([]int{5, 3, 4, 6, 7, 8, 9, 1}))
	assert.Equal(t, false, CheckForZero([]int{5, 3, 0, 6, 7, 8, 9, 1, 2}))
	assert.Equal(t, true, CheckForZero([]int{5, 3, 4, 6, 7, 8, 9, 1, 2}))
	assert.Equal(t, true, CheckForZero([]int{5, 3, 4, 6, 7, 8, 9, 1, 2}))
}

func TestGetRow(t *testing.T) {
	assert.Equal(t, []int{5, 3, 4, 6, 7, 8, 9, 1, 2}, GetRow(puzzle, 0))
	assert.Equal(t, []int{6, 7, 2, 1, 9, 5, 3, 4, 8}, GetRow(puzzle, 1))
	assert.Equal(t, []int{1, 9, 8, 3, 4, 2, 5, 6, 7}, GetRow(puzzle, 2))
	assert.Equal(t, []int{8, 5, 9, 7, 6, 1, 4, 2, 3}, GetRow(puzzle, 3))
	assert.Equal(t, []int{4, 2, 6, 8, 5, 3, 7, 9, 1}, GetRow(puzzle, 4))
	assert.Equal(t, []int{7, 1, 3, 9, 2, 4, 8, 5, 6}, GetRow(puzzle, 5))
	assert.Equal(t, []int{9, 6, 1, 5, 3, 7, 2, 8, 4}, GetRow(puzzle, 6))
	assert.Equal(t, []int{2, 8, 7, 4, 1, 9, 6, 3, 5}, GetRow(puzzle, 7))
	assert.Equal(t, []int{3, 4, 5, 2, 8, 6, 1, 7, 9}, GetRow(puzzle, 8))
}
