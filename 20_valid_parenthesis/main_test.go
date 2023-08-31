package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {

	case1 := "()"
	assert.True(t, IsValid(case1))

	case2 := "()[]{}"
	assert.True(t, IsValid(case2))

	case3 := "(]"
	assert.False(t, IsValid(case3))

	case4 := "([)]"
	assert.False(t, IsValid(case4))

	case5 := "{[]}"
	assert.True(t, IsValid(case5))

	case6 := "["
	assert.False(t, IsValid(case6))
}
