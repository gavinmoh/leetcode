package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyHashMap(t *testing.T) {
	obj := Constructor()
	obj.Put(1, 1)
	obj.Put(2, 2)
	assert.Equal(t, 1, obj.Get(1))
	assert.Equal(t, -1, obj.Get(3))
	obj.Put(2, 1)
	assert.Equal(t, 1, obj.Get(2))
	obj.Remove(2)
	assert.Equal(t, -1, obj.Get(2))
}
