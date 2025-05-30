package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClosestMeetingNode(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		edges := []int{2, 2, 3, -1}
		node1 := 0
		node2 := 1
		expected := 2

		assert.Equal(t, expected, closestMeetingNode(edges, node1, node2))
	})

	t.Run("test case 2", func(t *testing.T) {
		edges := []int{1, 2, -1}
		node1 := 0
		node2 := 2
		expected := 2

		assert.Equal(t, expected, closestMeetingNode(edges, node1, node2))
	})
}
