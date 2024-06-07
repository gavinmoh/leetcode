package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReplaceWords(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		dictionary := []string{"cat", "bat", "rat"}
		sentence := "the cattle was rattled by the battery"
		expected := "the cat was rat by the bat"

		assert.Equal(t, expected, replaceWords(dictionary, sentence))
	})

	t.Run("test case 2", func(t *testing.T) {
		dictionary := []string{"a", "b", "c"}
		sentence := "aadsfasf absbs bbab cadsfafs"
		expected := "a a b c"

		assert.Equal(t, expected, replaceWords(dictionary, sentence))
	})

	t.Run("test case 3", func(t *testing.T) {
		dictionary := []string{"a", "aa", "aaa", "aaaa"}
		sentence := "a aa a aaaa aaa aaa aaa aaaaaa bbb baba ababa"
		expected := "a a a a a a a a bbb baba a"

		assert.Equal(t, expected, replaceWords(dictionary, sentence))
	})
}
