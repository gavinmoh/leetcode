package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumTotalDistance(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		robot := []int{0, 4, 6}
		factory := [][]int{{2, 2}, {6, 2}}
		expected := int64(4)

		assert.Equal(t, expected, minimumTotalDistance(robot, factory))
	})

	t.Run("test case 2", func(t *testing.T) {
		robot := []int{1, -1}
		factory := [][]int{{-2, 1}, {2, 1}}
		expected := int64(2)

		assert.Equal(t, expected, minimumTotalDistance(robot, factory))
	})
}
