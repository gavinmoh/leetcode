package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThreeSum(t *testing.T) {
	case1 := []int{-1, 0, 1, 2, -1, -4}
	case1Expected := [][]int{{-1, -1, 2}, {-1, 0, 1}}
	assert.Equal(t, len(case1Expected), len(ThreeSum(case1)))
	for _, expected := range case1Expected {
		assert.Contains(t, ThreeSum(case1), expected)
	}

	case2 := []int{0, 1, 1}
	case2Expected := [][]int{}
	assert.Equal(t, len(case2Expected), len(ThreeSum(case2)))
	for _, expected := range case2Expected {
		assert.Contains(t, ThreeSum(case2), expected)
	}

	case3 := []int{0, 0, 0}
	case3Expected := [][]int{{0, 0, 0}}
	assert.Equal(t, len(case3Expected), len(ThreeSum(case3)))
	for _, expected := range case3Expected {
		assert.Contains(t, ThreeSum(case3), expected)
	}
}
