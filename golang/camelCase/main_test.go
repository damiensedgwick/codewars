package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var str, slice = "camel case string", []string{
	"camel",
	"case",
	"string",
}

func TestMakeStringSlice(t *testing.T) {
	assert.Equal(t, slice, MakeStringSlice(str))
}

var sliceWithCaps = []string{
	"Camel",
	"Case",
	"String",
}

func TestCapitaliseEachStringInSlice(t *testing.T) {
	assert.Equal(t, sliceWithCaps, CapitaliseEachStringInSlice(slice))
}

var camelCaseStr = "CamelCaseString"

func TestMakeCamelCaseString(t *testing.T) {
	assert.Equal(t, camelCaseStr, MakeCamelCaseString(sliceWithCaps))
}
