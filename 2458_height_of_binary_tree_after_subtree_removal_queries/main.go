package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type LevelInfo struct {
	MaxHeight    int
	SecondHeight int
}

func treeQueries(root *TreeNode, queries []int) []int {
	maxHeights := map[int]*LevelInfo{}
	levels := map[int]int{}  // level of each node
	heights := map[int]int{} // height of each tree

	var dfs func(node *TreeNode, level int) int
	dfs = func(node *TreeNode, level int) int {
		levels[node.Val] = level

		height := 0
		if node.Left != nil {
			height = max(height, dfs(node.Left, level+1)+1)
		}

		if node.Right != nil {
			height = max(height, dfs(node.Right, level+1)+1)
		}

		heights[node.Val] = height
		levelInfo := maxHeights[level]

		if levelInfo == nil {
			maxHeights[level] = &LevelInfo{
				MaxHeight:    height,
				SecondHeight: -1,
			}
		} else {
			if height > levelInfo.MaxHeight {
				levelInfo.SecondHeight = levelInfo.MaxHeight
				levelInfo.MaxHeight = height
			} else if height > levelInfo.SecondHeight {
				levelInfo.SecondHeight = height
			}
		}

		return height
	}

	dfs(root, 0)

	result := make([]int, len(queries))

	for i, node := range queries {
		level := levels[node]
		levelInfo := maxHeights[level]
		height := heights[node]

		if levelInfo.SecondHeight == -1 {
			result[i] = heights[root.Val] - height - 1
		} else if levelInfo.MaxHeight == height {
			result[i] = levelInfo.SecondHeight + level
		} else {
			result[i] = levelInfo.MaxHeight + level
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
