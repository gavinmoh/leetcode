package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSmallestInfiniteSet(t *testing.T) {
	set1 := Constructor()
	set1.AddBack(2)
	assert.Equal(t, 1, set1.PopSmallest())
	assert.Equal(t, 2, set1.PopSmallest())
	assert.Equal(t, 3, set1.PopSmallest())
	set1.AddBack(1)
	assert.Equal(t, 1, set1.PopSmallest())
	assert.Equal(t, 4, set1.PopSmallest())
	assert.Equal(t, 5, set1.PopSmallest())

	set2 := Constructor()
	assert.Equal(t, 1, set2.PopSmallest())
	set2.AddBack(1)
	assert.Equal(t, 1, set2.PopSmallest())
	assert.Equal(t, 2, set2.PopSmallest())
	assert.Equal(t, 3, set2.PopSmallest())
	set2.AddBack(2)
	set2.AddBack(3)
	assert.Equal(t, 2, set2.PopSmallest())
	assert.Equal(t, 3, set2.PopSmallest())
}
