package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyCalendar(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		events := [][]int{{10, 20}, {15, 25}, {20, 30}}
		expected := []bool{true, false, true}

		calendar := Constructor()
		for i, event := range events {
			result := calendar.Book(event[0], event[1])
			assert.Equal(t, expected[i], result)
		}
	})

	t.Run("test case 2", func(t *testing.T) {
		events := [][]int{{47, 50}, {33, 41}, {39, 45}, {33, 42}, {25, 32}, {26, 35}, {19, 25}, {3, 8}, {8, 13}, {18, 27}}
		expected := []bool{true, true, false, false, true, false, true, true, true, false}

		calendar := Constructor()
		for i, event := range events {
			result := calendar.Book(event[0], event[1])
			assert.Equal(t, expected[i], result)
		}
	})

	t.Run("test case 3", func(t *testing.T) {
		events := [][]int{{20, 29}, {13, 22}, {44, 50}, {1, 7}, {2, 10}, {14, 20}, {19, 25}, {36, 42}, {45, 50}, {47, 50}, {39, 45}, {44, 50}, {16, 25}, {45, 50}, {45, 50}, {12, 20}, {21, 29}, {11, 20}, {12, 17}, {34, 40}, {10, 18}, {38, 44}, {23, 32}, {38, 44}, {15, 20}, {27, 33}, {34, 42}, {44, 50}, {35, 40}, {24, 31}}
		expected := []bool{true, false, true, true, false, true, false, true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false}

		calendar := Constructor()
		for i, event := range events {
			result := calendar.Book(event[0], event[1])
			assert.Equal(t, expected[i], result)
		}
	})
}
