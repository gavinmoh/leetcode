package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRow(t *testing.T) {
	case1Expected := []int{1, 3, 3, 1}
	assert.Equal(t, case1Expected, getRow(3))

	case2Expected := []int{1}
	assert.Equal(t, case2Expected, getRow(0))

	case3Expected := []int{1, 1}
	assert.Equal(t, case3Expected, getRow(1))
}
