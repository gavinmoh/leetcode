package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyAtoi(t *testing.T) {
	case1 := "-+12"
	assert.Equal(t, 0, myAtoi(case1))

	case2 := "   -42"
	assert.Equal(t, -42, myAtoi(case2))

	case3 := "9223372036854775808"
	assert.Equal(t, 2147483647, myAtoi(case3))

	case4 := "  0000000000012345678"
	assert.Equal(t, 12345678, myAtoi(case4))
}
