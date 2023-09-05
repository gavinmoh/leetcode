package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThreeSumClosest(t *testing.T) {

	// case1Nums := []int{-1, 2, 1, -4}
	// case1Target := 1
	// assert.Equal(t, 2, ThreeSumClosest(case1Nums, case1Target))

	// case2Nums := []int{0, 0, 0}
	// case2Target := 1
	// assert.Equal(t, 0, ThreeSumClosest(case2Nums, case2Target))

	// case3Nums := []int{1, 1, 1, 0}
	// case3Target := -100
	// assert.Equal(t, 2, ThreeSumClosest(case3Nums, case3Target))

	case4Nums := []int{0, 3, 97, 102, 200}
	case4Target := 300
	assert.Equal(t, 300, ThreeSumClosest(case4Nums, case4Target))
}
