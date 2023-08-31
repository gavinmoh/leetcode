package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestCommonPrefix(t *testing.T) {

	case1 := []string{"flower", "flow", "flight"}
	assert.Equal(t, "fl", LongestCommonPrefix(case1))

	case2 := []string{"dog", "racecar", "car"}
	assert.Equal(t, "", LongestCommonPrefix(case2))

	case3 := []string{""}
	assert.Equal(t, "", LongestCommonPrefix(case3))

	case4 := []string{"a"}
	assert.Equal(t, "a", LongestCommonPrefix(case4))

	case5 := []string{"", ""}
	assert.Equal(t, "", LongestCommonPrefix(case5))

	case6 := []string{"ab", "a"}
	assert.Equal(t, "a", LongestCommonPrefix(case6))
}
