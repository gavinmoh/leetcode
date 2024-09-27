package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyCalendarTwo(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		events := [][]int{{10, 20}, {50, 60}, {10, 40}, {5, 15}, {5, 10}, {25, 55}}
		expected := []bool{true, true, true, false, true, true}

		calendar := Constructor()
		for i, event := range events {
			assert.Equal(t, expected[i], calendar.Book(event[0], event[1]))
		}
	})
}
