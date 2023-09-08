package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	case1Expected := [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}
	case1Actual := Generate(5)
	assert.Equal(t, case1Expected, case1Actual)

	case2Expected := [][]int{{1}}
	case2Actual := Generate(1)
	assert.Equal(t, case2Expected, case2Actual)

}
