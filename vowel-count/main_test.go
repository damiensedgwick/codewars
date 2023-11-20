package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	vowels := []string{"a", "e", "i", "o", "u"}

	assert.Equal(t, true, Contains(vowels, "a"))
	assert.Equal(t, false, Contains(vowels, "x"))
}

func TestGetCount(t *testing.T) {
	assert.Equal(t, 5, GetCount("abracadabra"))
}
