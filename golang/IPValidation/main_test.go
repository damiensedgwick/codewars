package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateInput(t *testing.T) {
	validInputs := []string{
		"1.2.3.4",
		"123.45.67.89",
	}
	assert.Equal(t, true, ValidateInput(validInputs[0]))
	assert.Equal(t, true, ValidateInput(validInputs[1]))

	invalidInputs := []string{
		"1.2.3",
		"1.2.3.4.5",
		"123.456.78.90",
		"123.045.067.089",
		"n.w.r.q",
		"t.d.h.v",
		"b.w.w.v",
	}
	assert.Equal(t, false, ValidateInput(invalidInputs[0]))
	assert.Equal(t, false, ValidateInput(invalidInputs[1]))
	assert.Equal(t, false, ValidateInput(invalidInputs[2]))
	assert.Equal(t, false, ValidateInput(invalidInputs[3]))
	assert.Equal(t, false, ValidateInput(invalidInputs[4]))
	assert.Equal(t, false, ValidateInput(invalidInputs[5]))
	assert.Equal(t, false, ValidateInput(invalidInputs[6]))
}
