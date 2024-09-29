package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAllOne(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		obj := Constructor()
		obj.Inc("hello")
		obj.Inc("hello")
		assert.Equal(t, "hello", obj.GetMaxKey())
		assert.Equal(t, "hello", obj.GetMinKey())
		obj.Inc("leet")
		assert.Equal(t, "leet", obj.GetMinKey())
	})

	t.Run("test case 2", func(t *testing.T) {
		obj := Constructor()
		obj.Inc("a")
		obj.Inc("b")
		obj.Inc("b")
		obj.Inc("c")
		obj.Inc("c")
		obj.Inc("c")
		obj.Dec("b")
		obj.Dec("b")
		assert.Equal(t, "a", obj.GetMinKey())
		obj.Dec("a")
		assert.Equal(t, "c", obj.GetMaxKey())
		assert.Equal(t, "c", obj.GetMinKey())
	})

	t.Run("test case 3", func(t *testing.T) {
		obj := Constructor()
		obj.Inc("a")
		obj.Inc("b")
		obj.Inc("b")
		obj.Inc("b")
		obj.Inc("b")
		obj.Dec("b")
		obj.Dec("b")
		assert.Equal(t, "b", obj.GetMaxKey())
		assert.Equal(t, "a", obj.GetMinKey())
	})
}
