package main

import "strconv"

type TrieNode struct {
	Children    [10]*TrieNode
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
		idx := char - '0'
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

func (this *Trie) StartsWith(prefix string) bool {
	current := this.Root
	for _, char := range prefix {
		idx := char - '0'
		if current.Children[idx] == nil {
			return false
		}
		current = current.Children[idx]
	}
	return true
}

func longestCommonPrefix(arr1 []int, arr2 []int) int {
	trie := Constructor()
	for _, num := range arr1 {
		str := strconv.Itoa(num)
		trie.Insert(str)
	}

	longest := 0
	for _, num := range arr2 {
		str := strconv.Itoa(num)
		prefixLength := len(str)
		for prefixLength > 0 {
			if trie.StartsWith(str[0:prefixLength]) {
				break
			}
			prefixLength--
		}
		if prefixLength > longest {
			longest = prefixLength
		}
	}

	return longest
}
