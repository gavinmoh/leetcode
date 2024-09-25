package main

type TrieNode struct {
	Children    [26]*TrieNode
	IsEndOfWord bool
	Counter     int
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
		current.Counter++
	}
	current.IsEndOfWord = true
}

func (this *Trie) SumScores(word string) int {
	current := this.Root
	score := 0
	for _, char := range word {
		idx := char - 'a'
		if current.Children[idx] == nil {
			return score
		}
		current = current.Children[idx]
		score += current.Counter
	}
	return score
}

func sumPrefixScores(words []string) []int {
	trie := Constructor()
	for _, word := range words {
		trie.Insert(word)
	}

	result := []int{}
	for _, word := range words {
		result = append(result, trie.SumScores(word))
	}

	return result
}
