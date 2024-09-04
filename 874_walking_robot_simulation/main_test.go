package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRobotSim(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		commands := []int{4, -1, 3}
		obstacles := [][]int{}
		expected := 25

		assert.Equal(t, expected, robotSim(commands, obstacles))
	})

	t.Run("test case 2", func(t *testing.T) {
		commands := []int{4, -1, 4, -2, 4}
		obstacles := [][]int{{2, 4}}
		expected := 65

		assert.Equal(t, expected, robotSim(commands, obstacles))
	})

	t.Run("test case 3", func(t *testing.T) {
		commands := []int{6, -1, -1, 6}
		obstacles := [][]int{}
		expected := 36

		assert.Equal(t, expected, robotSim(commands, obstacles))
	})
}
