package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumSubarrayMins(t *testing.T) {
	t.Run("should return 17", func(t *testing.T) {
		arr := []int{3, 1, 2, 4}
		expected := 17
		assert.Equal(t, expected, sumSubarrayMins(arr))
	})

	t.Run("should return 444", func(t *testing.T) {
		arr := []int{11, 81, 94, 43, 3}
		expected := 444
		assert.Equal(t, expected, sumSubarrayMins(arr))
	})

}
