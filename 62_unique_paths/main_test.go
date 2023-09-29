package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniquePaths(t *testing.T) {
	case1m := 3
	case1n := 7
	assert.Equal(t, 28, uniquePaths(case1m, case1n))

	case2m := 3
	case2n := 2
	assert.Equal(t, 3, uniquePaths(case2m, case2n))
}
