package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPowerOfFour(t *testing.T) {
	assert.True(t, isPowerOfFour(16))
	assert.False(t, isPowerOfFour(5))
	assert.True(t, isPowerOfFour(1))
}
