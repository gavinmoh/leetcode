package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createBinaryTree(descriptions [][]int) *TreeNode {
	nodes := make(map[int]*TreeNode)
	children := make(map[int]bool)

	for _, description := range descriptions {
		parent := description[0]
		child := description[1]
		isLeft := (description[2] == 1)
		children[child] = true

		if _, ok := nodes[child]; !ok {
			newNode := &TreeNode{
				Val: child,
			}
			nodes[child] = newNode
		}

		if _, ok := nodes[parent]; !ok {
			newNode := &TreeNode{
				Val: parent,
			}
			nodes[parent] = newNode
		}

		if isLeft {
			nodes[parent].Left = nodes[child]
		} else {
			nodes[parent].Right = nodes[child]
		}
	}

	var root *TreeNode
	for node := range nodes {
		if _, ok := children[node]; !ok {
			root = nodes[node]
			break
		}
	}

	return root
}
