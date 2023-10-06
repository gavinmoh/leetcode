package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntegerBreak(t *testing.T) {
	case1 := 2
	assert.Equal(t, 1, integerBreak(case1))

	case2 := 10
	assert.Equal(t, 36, integerBreak(case2))
}
