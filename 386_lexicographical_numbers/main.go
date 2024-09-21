package main

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

func (trie *Trie) Insert(number int) {
	current := trie.Root
	// split number into digits
	digits := []int{}
	for number > 0 {
		digit := number % 10
		digits = append(digits, digit)
		number = number / 10
	}
	// iterating from backwards
	for i := len(digits) - 1; i >= 0; i-- {
		idx := digits[i]
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
		trie.Insert(i)
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
