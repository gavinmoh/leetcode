package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumOfArrays(t *testing.T) {
	case1 := numOfArrays(2, 3, 1)
	assert.Equal(t, 6, case1)

	case2 := numOfArrays(5, 2, 3)
	assert.Equal(t, 0, case2)

	case3 := numOfArrays(9, 1, 1)
	assert.Equal(t, 1, case3)
}
