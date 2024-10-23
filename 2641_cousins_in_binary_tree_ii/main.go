package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func replaceValueInTree(root *TreeNode) *TreeNode {
	nodesMap := map[int][]*TreeNode{}

	// bfs
	nodes := []*TreeNode{root}
	level := 0
	for len(nodes) > 0 {
		nextNodes := []*TreeNode{}
		if _, ok := nodesMap[level]; !ok {
			nodesMap[level] = []*TreeNode{}
		}
		for _, node := range nodes {
			nodesMap[level] = append(nodesMap[level], node)

			if node.Left != nil {
				nextNodes = append(nextNodes, node.Left)
			}

			if node.Right != nil {
				nextNodes = append(nextNodes, node.Right)
			}
		}
		nodes = nextNodes
		level++
	}

	root.Val = 0
	for _, siblings := range nodesMap {
		totalSum := 0
		siblingSum := map[*TreeNode]int{}
		for _, sibling := range siblings {
			sum := 0
			if sibling.Left != nil {
				sum += sibling.Left.Val
			}
			if sibling.Right != nil {
				sum += sibling.Right.Val
			}
			siblingSum[sibling] = sum
			totalSum += sum
		}
		for _, sibling := range siblings {
			if sibling.Left != nil {
				sibling.Left.Val = totalSum - siblingSum[sibling]
			}
			if sibling.Right != nil {
				sibling.Right.Val = totalSum - siblingSum[sibling]
			}
		}
	}

	return root
}
