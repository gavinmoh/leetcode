package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	str := "ohvhjdml"
	assert.Equal(t, 6, lengthOfLongestSubstring(str))
}
