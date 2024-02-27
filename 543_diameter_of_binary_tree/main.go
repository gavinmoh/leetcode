package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	cache := make(map[*TreeNode]int)
	var maxDepth func(*TreeNode) int
	maxDepth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		if cachedResult, ok := cache[node]; ok {
			return cachedResult
		}

		left := maxDepth(node.Left)
		right := maxDepth(node.Right)

		if left > right {
			cache[node] = left + 1
		} else {
			cache[node] = right + 1
		}

		return cache[node]
	}

	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		left := maxDepth(node.Left)
		right := maxDepth(node.Right)

		diameter := left + right

		return max(diameter, max(dfs(node.Left), dfs(node.Right)))
	}

	return dfs(root)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
