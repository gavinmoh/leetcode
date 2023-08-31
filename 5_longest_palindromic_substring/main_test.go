package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestPalindrome(t *testing.T) {
	str := "babad"
	assert.Equal(t, "bab", longestPalindrome(str))

}
