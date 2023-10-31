package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindArray(t *testing.T) {
	t.Run("should return [5,7,2,3,2]", func(t *testing.T) {
		pref := []int{5, 2, 0, 3, 1}
		expected := []int{5, 7, 2, 3, 2}
		assert.Equal(t, expected, findArray(pref))
	})

	t.Run("should return [13]", func(t *testing.T) {
		pref := []int{13}
		expected := []int{13}
		assert.Equal(t, expected, findArray(pref))
	})
}
