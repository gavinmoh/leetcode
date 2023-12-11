package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindSpecialInteger(t *testing.T) {
	t.Run("should return 6", func(t *testing.T) {
		arr := []int{1, 2, 2, 6, 6, 6, 6, 7, 10}
		expected := 6
		assert.Equal(t, expected, findSpecialInteger(arr))
	})

	t.Run("should return 1", func(t *testing.T) {
		arr := []int{1, 1}
		expected := 1
		assert.Equal(t, expected, findSpecialInteger(arr))
	})
}
