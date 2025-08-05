package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumOfUnplacedFruits(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		fruits := []int{4, 2, 5}
		baskets := []int{3, 5, 4}
		expected := 1

		assert.Equal(t, expected, numOfUnplacedFruits(fruits, baskets))
	})

	t.Run("test case 2", func(t *testing.T) {
		fruits := []int{3, 6, 1}
		baskets := []int{6, 4, 7}
		expected := 0

		assert.Equal(t, expected, numOfUnplacedFruits(fruits, baskets))
	})
}
