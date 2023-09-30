package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFind132Pattern(t *testing.T) {
	case1 := []int{1, 2, 3, 4}
	assert.False(t, find132pattern(case1))

	case2 := []int{3, 1, 4, 2}
	assert.True(t, find132pattern(case2))

	case3 := []int{-1, 3, 2, 0}
	assert.True(t, find132pattern(case3))
}
