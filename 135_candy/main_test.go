package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	case1 := []int{1, 0, 2}
	assert.Equal(t, 5, Candy(case1))

	case2 := []int{1, 2, 2}
	assert.Equal(t, 4, Candy(case2))
}
