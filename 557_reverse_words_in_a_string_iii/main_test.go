package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseWords(t *testing.T) {
	case1 := "Let's take LeetCode contest"
	assert.Equal(t, "s'teL ekat edoCteeL tsetnoc", reverseWords(case1))

	case2 := "God Ding"
	assert.Equal(t, "doG gniD", reverseWords(case2))
}
