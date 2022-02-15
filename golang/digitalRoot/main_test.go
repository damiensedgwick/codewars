package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMakeSlice(t *testing.T) {
	assert.Equal(t, []int{2, 3}, MakeSlice(23))
	assert.Equal(t, []int{5, 6, 7}, MakeSlice(567))
	assert.Equal(t, []int{9, 8, 7, 6, 5}, MakeSlice(98765))
}

func TestSumSliceDigits(t *testing.T) {
	assert.Equal(t, 7, SumSliceDigits([]int{1, 6}))
	assert.Equal(t, 6, SumSliceDigits([]int{9, 4, 2}))
	assert.Equal(t, 6, SumSliceDigits([]int{1, 3, 2, 1, 8, 9}))
	assert.Equal(t, 2, SumSliceDigits([]int{4, 9, 3, 1, 9, 3}))
}
