package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChampagneTower(t *testing.T) {
	case1 := champagneTower(1, 1, 1)
	assert.Equal(t, 0.0, case1)

	case2 := champagneTower(2, 1, 1)
	assert.Equal(t, 0.5, case2)

	case3 := champagneTower(100000009, 33, 17)
	assert.Equal(t, 1.0, case3)

	case4 := champagneTower(25, 6, 1)
	assert.Equal(t, 0.1875, case4)
}
