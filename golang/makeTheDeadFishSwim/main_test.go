package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIncrementOne(t *testing.T) {
	assert.Equal(t, 1, IncrementOne(0))
	assert.Equal(t, 5, IncrementOne(4))
	assert.Equal(t, 100, IncrementOne(99))
}

func TestDecrementOne(t *testing.T) {
	assert.Equal(t, 0, DecrementOne(1))
	assert.Equal(t, 4, DecrementOne(5))
	assert.Equal(t, 99, DecrementOne(100))
}

func TestSquare(t *testing.T) {
	assert.Equal(t, 1, Square(1))
	assert.Equal(t, 25, Square(5))
	assert.Equal(t, 64, Square(8))
}

func TestSplitString(t *testing.T) {
	assert.Equal(t, []string{
		"i",
		"i",
		"i",
		"s",
		"d",
		"o",
		"s",
		"o",
	}, SplitString("iiisdoso"))
}

func TestParse(t *testing.T) {
	assert.Equal(t, []int{8, 64}, Parse("iiisdoso"))
}
