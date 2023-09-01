package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxArea(t *testing.T) {
	case1 := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	assert.Equal(t, 49, MaxArea(case1))

	case2 := []int{1, 1}
	assert.Equal(t, 1, MaxArea(case2))
}
