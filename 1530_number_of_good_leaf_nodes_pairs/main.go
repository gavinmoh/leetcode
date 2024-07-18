package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countPairs(root *TreeNode, distance int) int {
	pairs := 0

	var dfs func(node *TreeNode) []int
	dfs = func(node *TreeNode) []int {
		if node == nil {
			return []int{}
		}

		if node.Left == nil && node.Right == nil {
			return []int{1}
		}
		left := dfs(node.Left)
		right := dfs(node.Right)

		for _, leftDistance := range left {
			for _, rightDistance := range right {
				if (leftDistance + rightDistance) <= distance {
					pairs++
				}
			}
		}

		result := []int{}
		for _, leftDistance := range left {
			result = append(result, leftDistance+1)
		}

		for _, rightDistance := range right {
			result = append(result, rightDistance+1)
		}

		return result
	}

	dfs(root)

	return pairs
}
