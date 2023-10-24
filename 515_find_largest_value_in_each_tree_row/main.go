package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	// edge case
	if root == nil {
		return []int{}
	}

	minInt := -1 << 31
	maxLevel := 0

	rows := make(map[int]int, 10001)
	for i := 0; i < 10001; i++ {
		rows[i] = minInt
	}

	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, level int) {
		if node == nil {
			return
		}

		if level > maxLevel {
			maxLevel = level
		}

		if node.Val > rows[level] {
			rows[level] = node.Val
		}

		dfs(node.Left, level+1)
		dfs(node.Right, level+1)
	}

	dfs(root, 0)

	var result []int
	for i := 0; i <= maxLevel; i++ {
		result = append(result, rows[i])
	}

	return result
}
