package main

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthLargestLevelSum(root *TreeNode, k int) int64 {
	nodes := []*TreeNode{root}
	sums := []int64{}

	for len(nodes) > 0 {
		children := []*TreeNode{}
		sum := int64(0)
		for _, node := range nodes {
			sum += int64(node.Val)
			if node.Left != nil {
				children = append(children, node.Left)
			}
			if node.Right != nil {
				children = append(children, node.Right)
			}
		}
		sums = append(sums, sum)
		nodes = children
	}

	if k > len(sums) {
		return -1
	}

	sort.SliceStable(sums, func(i, j int) bool {
		return sums[i] > sums[j]
	})

	return sums[k-1]
}
