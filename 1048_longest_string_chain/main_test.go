package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestStrChain(t *testing.T) {
	case1 := []string{"a", "b", "ba", "bca", "bda", "bdca"}
	assert.Equal(t, 4, longestStrChain(case1))

	case2 := []string{"xbc", "pcxbcf", "xb", "cxbc", "pcxbc"}
	assert.Equal(t, 5, longestStrChain(case2))

	case3 := []string{"abcd", "dbqca"}
	assert.Equal(t, 1, longestStrChain(case3))
}
