package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFullBloomFlowers(t *testing.T) {
	case1Flowers := [][]int{{1, 6}, {3, 7}, {9, 12}, {4, 13}}
	case1People := []int{2, 3, 7, 11}
	assert.Equal(t, []int{1, 2, 2, 2}, fullBloomFlowers(case1Flowers, case1People))

	case2Flowers := [][]int{{1, 10}, {3, 3}}
	case2People := []int{3, 3, 2}
	assert.Equal(t, []int{2, 2, 1}, fullBloomFlowers(case2Flowers, case2People))
}
