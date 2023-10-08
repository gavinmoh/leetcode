package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxDotProduct(t *testing.T) {
	case1nums1 := []int{2, 1, -2, 5}
	case1nums2 := []int{3, 0, -6}
	assert.Equal(t, 18, maxDotProduct(case1nums1, case1nums2))

	case2nums1 := []int{3, -2}
	case2nums2 := []int{2, -6, 7}
	assert.Equal(t, 21, maxDotProduct(case2nums1, case2nums2))

	case3nums1 := []int{-1, -1}
	case3nums2 := []int{1, 1}
	assert.Equal(t, -1, maxDotProduct(case3nums1, case3nums2))
}
