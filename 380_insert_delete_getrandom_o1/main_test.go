package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomizedSet(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		obj := Constructor()
		assert.True(t, obj.Insert(1))
		assert.False(t, obj.Remove(2))
		assert.True(t, obj.Insert(2))
		firstRandom := obj.GetRandom()
		assert.Greater(t, firstRandom, 0)
		assert.Less(t, firstRandom, 3)
		assert.True(t, obj.Remove(1))
		assert.False(t, obj.Insert(2))
		assert.Equal(t, 2, obj.GetRandom()) // this can only return 2
	})
}
