package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMakeByteSlice(t *testing.T) {
	assert.Equal(t, []byte{
		116,
		101,
		115,
		116,
	}, MakeByteSlice("test"))
	assert.Equal(t, []byte{
		72,
		101,
		108,
		108,
		111,
		32,
		87,
		111,
		114,
		108,
		100,
		33,
	}, MakeByteSlice("Hello World!"))
}

// Encode should take a []byte and return a basE91 encoding string.
func TestEncode(t *testing.T) {
	assert.Equal(t, "fPNKd", Encode([]byte{
		116,
		101,
		115,
		116,
	}))
	assert.Equal(t, ">OwJh>Io0Tv!8PE", Encode([]byte{
		72,
		101,
		108,
		108,
		111,
		32,
		87,
		111,
		114,
		108,
		100,
		33,
	}))
}
