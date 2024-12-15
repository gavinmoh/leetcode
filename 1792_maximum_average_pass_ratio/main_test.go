package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxAverageRatio(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		classes := [][]int{{1, 2}, {3, 5}, {2, 2}}
		extraStudents := 2
		expected := 0.7833333333333333

		assert.Equal(t, expected, maxAverageRatio(classes, extraStudents))
	})

	t.Run("test case 2", func(t *testing.T) {
		classes := [][]int{{2, 4}, {3, 9}, {4, 5}, {2, 10}}
		extraStudents := 4
		expected := 0.5348484848484849

		assert.Equal(t, expected, maxAverageRatio(classes, extraStudents))
	})
}
