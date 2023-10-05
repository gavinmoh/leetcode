package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMajorityElement(t *testing.T) {
	case1 := []int{3, 2, 3}
	assert.Equal(t, []int{3}, majorityElement(case1))

	case2 := []int{1}
	assert.Equal(t, []int{1}, majorityElement(case2))

	case3 := []int{1, 2}
	assert.Equal(t, []int{1, 2}, majorityElement(case3))
}
