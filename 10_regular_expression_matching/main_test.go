package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsMatch(t *testing.T) {
	case1 := isMatch("aa", "a")
	assert.False(t, case1)

	case2 := isMatch("aa", "a*")
	assert.True(t, case2)

	case3 := isMatch("ab", ".*")
	assert.True(t, case3)

	case4 := isMatch("aab", "c*a*b")
	assert.True(t, case4)

	case5 := isMatch("ab", ".*c")
	assert.False(t, case5)

	case6 := isMatch("aaa", "aaaa")
	assert.False(t, case6)

	case7 := isMatch("aaa", "a*a")
	assert.True(t, case7)

	case8 := isMatch("ab", ".*..")
	assert.True(t, case8)

	case9 := isMatch("a", ".*..a*")
	assert.False(t, case9)

	case10 := isMatch("abbbcd", "ab*bbbcd")
	assert.True(t, case10)
}
