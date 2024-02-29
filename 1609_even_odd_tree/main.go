package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isEvenOddTree(root *TreeNode) bool {
	nodes := []*TreeNode{root}
	level := 0

	// bfs
	for len(nodes) > 0 {
		isEvenLevel := (level%2 == 0)
		nextNodes := []*TreeNode{}
		for i, node := range nodes {
			if node.Left != nil {
				nextNodes = append(nextNodes, node.Left)
			}

			if node.Right != nil {
				nextNodes = append(nextNodes, node.Right)
			}

			if isEvenLevel {
				if node.Val%2 == 0 {
					return false
				}
				if i > 0 && node.Val <= nodes[i-1].Val {
					return false
				}
			} else {
				if node.Val%2 != 0 {
					return false
				}
				if i > 0 && node.Val >= nodes[i-1].Val {
					return false
				}
			}
		}

		level += 1
		nodes = nextNodes
	}

	return true
}
