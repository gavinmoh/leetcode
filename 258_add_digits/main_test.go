package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddDigits(t *testing.T) {
	case1 := 38
	assert.Equal(t, 2, AddDigits(case1))

	case2 := 0
	assert.Equal(t, 0, AddDigits(case2))
}
