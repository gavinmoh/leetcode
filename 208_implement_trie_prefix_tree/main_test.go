package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrie(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		trie := Constructor()

		trie.Insert("apple")
		assert.True(t, trie.Search("apple"))
		assert.False(t, trie.Search("app"))
		assert.True(t, trie.StartsWith("app"))

		trie.Insert("app")
		assert.True(t, trie.Search("app"))
	})
}
