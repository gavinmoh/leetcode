package main

import "strings"

type TrieNode struct {
	Children    [26]*TrieNode
	IsEndOfWord bool
}

type Trie struct {
	Root *TrieNode
}

func Constructor() Trie {
	obj := Trie{
		Root: &TrieNode{},
	}
	return obj
}

func (this *Trie) Insert(word string) {
	current := this.Root
	for _, char := range word {
		idx := char - 'a'
		if current.Children[idx] != nil {
			current = current.Children[idx]
		} else {
			newNode := &TrieNode{}
			current.Children[idx] = newNode
			current = newNode
		}
	}
	current.IsEndOfWord = true
}

func (this *Trie) FindShortest(word string) string {
	current := this.Root
	for i, char := range word {
		idx := char - 'a'
		if current.Children[idx] == nil {
			break
		}
		current = current.Children[idx]
		if current.IsEndOfWord {
			return word[0 : i+1]
		}
	}
	return word
}

func replaceWords(dictionary []string, sentence string) string {
	trie := Constructor()
	for _, word := range dictionary {
		trie.Insert(word)
	}

	words := strings.Split(sentence, " ")
	output := make([]string, len(words))
	for idx, word := range words {
		output[idx] = trie.FindShortest(word)
	}

	return strings.Join(output, " ")
}
