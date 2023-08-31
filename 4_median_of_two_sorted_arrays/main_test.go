package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMedianSortedArrays(t *testing.T) {
	case1nums1 := []int{1, 3}
	case1nums2 := []int{2}
	assert.Equal(t, 2.00000, findMedianSortedArrays(case1nums1, case1nums2))

	case2nums1 := []int{1, 2}
	case2nums2 := []int{3, 4}
	assert.Equal(t, 2.50000, findMedianSortedArrays(case2nums1, case2nums2))
}
