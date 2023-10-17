package main

type Node struct {
	left  *Node
	right *Node
}

func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
	// build the tree
	nodes := make([]*Node, n)
	for i := 0; i < n; i++ {
		nodes[i] = &Node{}
	}

	for i := 0; i < n; i++ {
		if leftChild[i] != -1 {
			nodes[i].left = nodes[leftChild[i]]
		}

		if rightChild[i] != -1 {
			nodes[i].right = nodes[rightChild[i]]
		}
	}

	// check if nodes have more than one parent
	parents := make(map[*Node]int)
	for _, node := range nodes {
		if node.left != nil {
			parents[node.left]++
		}

		if node.right != nil {
			parents[node.right]++
		}
	}

	for _, count := range parents {
		if count > 1 {
			return false
		}
	}

	// check if there is a cycle
	visited := make(map[*Node]bool)
	var hasCycle func(node *Node) bool
	hasCycle = func(node *Node) bool {
		if node == nil {
			return false
		}

		if visited[node] {
			return true
		}

		visited[node] = true
		return hasCycle(node.left) || hasCycle(node.right)
	}

	if hasCycle(nodes[0]) {
		return false
	}

	// check if there is multiple roots
	var root *Node
	for _, node := range nodes {
		if parents[node] == 0 {
			if root != nil {
				return false
			}

			root = node
		}
	}

	// check if there is a root
	return root != nil

}
