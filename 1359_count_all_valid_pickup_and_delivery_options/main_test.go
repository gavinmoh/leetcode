package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountOrders(t *testing.T) {
	case1 := 1
	assert.Equal(t, 1, CountOrders(case1))

	case2 := 2
	assert.Equal(t, 6, CountOrders(case2))

	case3 := 3
	assert.Equal(t, 90, CountOrders(case3))
}
