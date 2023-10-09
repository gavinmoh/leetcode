package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchRange(t *testing.T) {
	case1nums := []int{5, 7, 7, 8, 8, 10}
	case1target := 8
	assert.Equal(t, []int{3, 4}, searchRange(case1nums, case1target))

	case2nums := []int{5, 7, 7, 8, 8, 10}
	case2target := 6
	assert.Equal(t, []int{-1, -1}, searchRange(case2nums, case2target))

	case3nums := []int{}
	case3target := 0
	assert.Equal(t, []int{-1, -1}, searchRange(case3nums, case3target))
}
