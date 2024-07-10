package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinOperations(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		logs := []string{"d1/", "d2/", "../", "d21/", "./"}
		expected := 2

		assert.Equal(t, expected, minOperations(logs))
	})

	t.Run("test case 2", func(t *testing.T) {
		logs := []string{"d1/", "d2/", "./", "d3/", "../", "d31/"}
		expected := 3

		assert.Equal(t, expected, minOperations(logs))
	})

	t.Run("test case 3", func(t *testing.T) {
		logs := []string{"d1/", "../", "../", "../"}
		expected := 0

		assert.Equal(t, expected, minOperations(logs))
	})
}
