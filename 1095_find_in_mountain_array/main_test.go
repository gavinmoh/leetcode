package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindInMountainArray(t *testing.T) {
	case1Array := MountainArray{data: []int{1, 2, 3, 4, 5, 3, 1}}
	case1Target := 3
	assert.Equal(t, 2, findInMountainArray(case1Target, &case1Array))

	case2Array := MountainArray{data: []int{0, 1, 2, 4, 2, 1}}
	case2Target := 3
	assert.Equal(t, -1, findInMountainArray(case2Target, &case2Array))
}
