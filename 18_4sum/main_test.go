package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFourSum(t *testing.T) {
	case1 := []int{1, 0, -1, 0, -2, 2}
	case1Target := 0
	case1Expected := [][]int{
		[]int{-2, -1, 1, 2},
		[]int{-2, 0, 0, 2},
		[]int{-1, 0, 0, 1},
	}
	assert.Equal(t, case1Expected, FourSum(case1, case1Target))

	case2 := []int{2, 2, 2, 2, 2}
	case2Target := 8
	case2Expected := [][]int{
		[]int{2, 2, 2, 2},
	}
	assert.Equal(t, case2Expected, FourSum(case2, case2Target))

	case3 := []int{-5, -4, -3, -2, -1, 0, 0, 1, 2, 3, 4, 5}
	case3Target := 0
	case3Expected := [][]int{
		{-5, -4, 4, 5},
		{-5, -3, 3, 5},
		{-5, -2, 2, 5},
		{-5, -2, 3, 4},
		{-5, -1, 1, 5},
		{-5, -1, 2, 4},
		{-5, 0, 0, 5},
		{-5, 0, 1, 4},
		{-5, 0, 2, 3},
		{-4, -3, 2, 5},
		{-4, -3, 3, 4},
		{-4, -2, 1, 5},
		{-4, -2, 2, 4},
		{-4, -1, 0, 5},
		{-4, -1, 1, 4},
		{-4, -1, 2, 3},
		{-4, 0, 0, 4},
		{-4, 0, 1, 3},
		{-3, -2, 0, 5},
		{-3, -2, 1, 4},
		{-3, -2, 2, 3},
		{-3, -1, 0, 4},
		{-3, -1, 1, 3},
		{-3, 0, 0, 3},
		{-3, 0, 1, 2},
		{-2, -1, 0, 3},
		{-2, -1, 1, 2},
		{-2, 0, 0, 2},
		{-1, 0, 0, 1},
	}
	assert.Equal(t, case3Expected, FourSum(case3, case3Target))
}
