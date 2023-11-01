package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) []int {
	freqMap := make(map[int]int) // val -> freq
	max := 0

	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		freqMap[node.Val]++
		if freqMap[node.Val] > max {
			max = freqMap[node.Val]
		}
		dfs(node.Left)
		dfs(node.Right)
	}

	dfs(root)

	var result []int
	for val, freq := range freqMap {
		if freq == max {
			result = append(result, val)
		}
	}

	return result

}
