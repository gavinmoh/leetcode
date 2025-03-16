package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepairCars(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		ranks := []int{4, 2, 3, 1}
		cars := 10
		expected := int64(16)

		assert.Equal(t, expected, repairCars(ranks, cars))
	})

	t.Run("test case 2", func(t *testing.T) {
		ranks := []int{5, 1, 8}
		cars := 6
		expected := int64(16)

		assert.Equal(t, expected, repairCars(ranks, cars))
	})
}
