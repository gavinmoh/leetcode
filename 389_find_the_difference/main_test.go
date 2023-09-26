package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindTheDifference(t *testing.T) {
	case1 := findTheDifference("abcd", "abcde")
	case1Expected := byte('e')
	assert.Equal(t, case1Expected, case1)

	case2 := findTheDifference("", "y")
	case2Expected := byte('y')
	assert.Equal(t, case2Expected, case2)
}
