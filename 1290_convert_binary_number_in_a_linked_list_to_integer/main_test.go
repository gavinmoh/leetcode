package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDecimalValue(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		head := &ListNode{Val: 1, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1}}}
		expected := 5

		assert.Equal(t, expected, getDecimalValue(head))
	})

	t.Run("test case 2", func(t *testing.T) {
		head := &ListNode{Val: 0}
		expected := 0

		assert.Equal(t, expected, getDecimalValue(head))
	})
}
