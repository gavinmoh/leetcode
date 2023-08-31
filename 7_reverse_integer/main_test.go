package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	x := 123
	assert.Equal(t, 321, reverse(x))
}
