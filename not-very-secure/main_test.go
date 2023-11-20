package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlphanumeric(t *testing.T) {
	assert.Equal(t, true, Alphanumeric("banana"))
	assert.Equal(t, false, Alphanumeric("ban ana"))
	assert.Equal(t, false, Alphanumeric("ban@ana"))
	assert.Equal(t, false, Alphanumeric("ban_ana"))
	assert.Equal(t, false, Alphanumeric(""))
}
