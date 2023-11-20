package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindUniq(t *testing.T) {
	var r1, r2, r3 float32 = 0.55, 2, 5
	assert.Equal(t, r1, FindUniq([]float32{0, 0, 0, 0, 0.55}))
	assert.Equal(t, r2, FindUniq([]float32{1, 2, 1, 1, 1}))
	assert.Equal(t, r3, FindUniq([]float32{4.4, 4.4, 4.4, 5, 4.4}))
}
