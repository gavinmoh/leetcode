package main

import (
	"strconv"
)

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

func lexicalOrder(n int) []int {
	trie := Constructor()
	for i := 1; i <= n; i++ {
		trie.Insert(strconv.Itoa(i))
	}
	result := []int{}

	var dfs func(*TrieNode, int, int)
	dfs = func(node *TrieNode, idx int, acc int) {
		if node == nil {
			return
		}

		num := (acc * 10) + idx

		if node.IsEndOfWord {
			result = append(result, num)
		}

		for i := 0; i < 10; i++ {
			dfs(node.Children[i], i, num)
		}
	}

	for i := 0; i < 10; i++ {
		dfs(trie.Root.Children[i], i, 0)
	}

	return result
}
