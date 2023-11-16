package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindDifferentBinaryString(t *testing.T) {
	t.Run("should return 00", func(t *testing.T) {
		nums := []string{"01", "10"}
		expected := []string{"00", "11"}
		assert.Contains(t, expected, findDifferentBinaryString(nums))
	})

	t.Run("should return 10", func(t *testing.T) {
		nums := []string{"00", "01"}
		expected := []string{"10", "11"}
		assert.Contains(t, expected, findDifferentBinaryString(nums))
	})

	t.Run("should return 000", func(t *testing.T) {
		nums := []string{"111", "011", "001"}
		expected := []string{"000", "010", "100", "101", "110"}
		assert.Contains(t, expected, findDifferentBinaryString(nums))
	})
}
