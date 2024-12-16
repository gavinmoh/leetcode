package main

import (
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minimumOperations(root *TreeNode) int {
	nodes := []*TreeNode{root}
	count := 0

	for level := 0; len(nodes) > 0; level++ {
		nextNodes := []*TreeNode{}
		values := []int{}
		for _, node := range nodes {
			if node == nil {
				continue
			}
			values = append(values, node.Val)
			nextNodes = append(nextNodes, node.Left, node.Right)
		}

		// clone and sort the array of values
		sorted := make([]int, len(values))
		copy(sorted, values)
		sort.Ints(sorted)

		valueIdx := map[int]int{}
		for i, value := range values {
			valueIdx[value] = i
		}

		for expectedIdx, currentValue := range sorted {
			actualIdx := valueIdx[currentValue]
			if actualIdx != expectedIdx {
				swapValue := values[expectedIdx]
				swapIdx := valueIdx[swapValue]
				// swap the mapping
				valueIdx[currentValue], valueIdx[swapValue] = swapIdx, actualIdx
				// swap the values
				values[actualIdx], values[swapIdx] = values[swapIdx], values[actualIdx]
				count++
			}
		}

		nodes = nextNodes
	}

	return count
}
