package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompareVersion(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		version1 := "1.01"
		version2 := "1.001"
		expected := 0

		assert.Equal(t, expected, compareVersion(version1, version2))
	})

	t.Run("test case 2", func(t *testing.T) {
		version1 := "1.0"
		version2 := "1.0.0"
		expected := 0

		assert.Equal(t, expected, compareVersion(version1, version2))
	})

	t.Run("test case 3", func(t *testing.T) {
		version1 := "0.1"
		version2 := "1.1"
		expected := -1

		assert.Equal(t, expected, compareVersion(version1, version2))
	})
}
