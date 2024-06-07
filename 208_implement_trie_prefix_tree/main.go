package main

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

func (this *Trie) Search(word string) bool {
	current := this.Root
	for _, char := range word {
		idx := char - 'a'
		if current.Children[idx] == nil {
			return false
		}
		current = current.Children[idx]
	}
	return current.IsEndOfWord
}

func (this *Trie) StartsWith(prefix string) bool {
	current := this.Root
	for _, char := range prefix {
		idx := char - 'a'
		if current.Children[idx] == nil {
			return false
		}
		current = current.Children[idx]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
