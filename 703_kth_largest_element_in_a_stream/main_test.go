package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKthLargest(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		k := 3
		nums := []int{4, 5, 8, 2}
		kthLargest := Constructor(k, nums)

		assert.Equal(t, 4, kthLargest.Add(3))
		assert.Equal(t, 5, kthLargest.Add(5))
		assert.Equal(t, 5, kthLargest.Add(10))
		assert.Equal(t, 8, kthLargest.Add(9))
		assert.Equal(t, 8, kthLargest.Add(4))
	})
}
