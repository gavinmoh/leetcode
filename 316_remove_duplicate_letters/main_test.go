package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicateLetters(t *testing.T) {
	// case1 := "bcabc"
	// assert.Equal(t, "abc", removeDuplicateLetters(case1))

	// case2 := "cbacdcbc"
	// assert.Equal(t, "acdb", removeDuplicateLetters(case2))

	case3 := "bbcaac"
	assert.Equal(t, "bac", removeDuplicateLetters(case3))

}
