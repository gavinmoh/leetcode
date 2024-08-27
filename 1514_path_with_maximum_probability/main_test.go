package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProbability(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		n := 3
		edges := [][]int{{0, 1}, {1, 2}, {0, 2}}
		succProb := []float64{0.5, 0.5, 0.2}
		start := 0
		end := 2
		expected := math.Pow(10, math.Log10(0.5)+math.Log10(0.5))

		assert.Equal(t, expected, maxProbability(n, edges, succProb, start, end))
	})

	t.Run("test case 2", func(t *testing.T) {
		n := 3
		edges := [][]int{{0, 1}, {1, 2}, {0, 2}}
		succProb := []float64{0.5, 0.5, 0.3}
		start := 0
		end := 2
		expected := math.Pow(10, math.Log10(0.3))

		assert.Equal(t, expected, maxProbability(n, edges, succProb, start, end))
	})

	t.Run("test case 3", func(t *testing.T) {
		n := 3
		edges := [][]int{{0, 1}}
		succProb := []float64{0.5}
		start := 0
		end := 2
		expected := float64(0)

		assert.Equal(t, expected, maxProbability(n, edges, succProb, start, end))
	})
}
