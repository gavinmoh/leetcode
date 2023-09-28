package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortArrayByParity(t *testing.T) {
	case1 := []int{3, 1, 2, 4}
	assert.Equal(t, []int{2, 4, 3, 1}, sortArrayByParity(case1))

	case2 := []int{0}
	assert.Equal(t, []int{0}, sortArrayByParity(case2))
}
