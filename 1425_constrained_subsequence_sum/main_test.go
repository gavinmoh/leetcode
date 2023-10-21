package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstrainedSubsetSum(t *testing.T) {
	case1Nums := []int{10, 2, -10, 5, 20}
	case1K := 2
	assert.Equal(t, 37, constrainedSubsetSum(case1Nums, case1K))

	case2Nums := []int{-1, -2, -3}
	case2K := 1
	assert.Equal(t, -1, constrainedSubsetSum(case2Nums, case2K))

	case3Nums := []int{10, -2, -10, -5, 20}
	case3K := 2
	assert.Equal(t, 23, constrainedSubsetSum(case3Nums, case3K))
}
