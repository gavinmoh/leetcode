package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindDuplicate(t *testing.T) {
	case1 := []int{1, 3, 4, 2, 2}
	assert.Equal(t, 2, findDuplicate(case1))

	case2 := []int{3, 1, 3, 4, 2}
	assert.Equal(t, 3, findDuplicate(case2))

	case3 := []int{2, 2, 2, 2}
	assert.Equal(t, 2, findDuplicate(case3))
}
