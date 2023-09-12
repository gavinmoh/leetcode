package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinDeletions(t *testing.T) {
	case1 := "aab"
	assert.Equal(t, 0, MinDeletions(case1))

	case2 := "aaabbbcc"
	assert.Equal(t, 2, MinDeletions(case2))

	case3 := "ceabaacb"
	assert.Equal(t, 2, MinDeletions(case3))

	case4 := "abcabc"
	assert.Equal(t, 3, MinDeletions(case4))

	case5 := "bbcebab"
	assert.Equal(t, 2, MinDeletions(case5))
}
