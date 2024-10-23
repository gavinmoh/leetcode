package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type NodeInfo struct {
	Parent *TreeNode
	Level  int
}

func isCousins(root *TreeNode, x int, y int) bool {
	nodesMap := map[int]*NodeInfo{}
	nodesMap[root.Val] = &NodeInfo{Level: 0}

	// bfs
	nodes := []*TreeNode{root}
	level := 1
	for len(nodes) > 0 {
		nextNodes := []*TreeNode{}
		for _, node := range nodes {
			if node.Left != nil {
				nodesMap[node.Left.Val] = &NodeInfo{Parent: node, Level: level}
				nextNodes = append(nextNodes, node.Left)
			}

			if node.Right != nil {
				nodesMap[node.Right.Val] = &NodeInfo{Parent: node, Level: level}
				nextNodes = append(nextNodes, node.Right)
			}
		}
		nodes = nextNodes
		level++

		if _, ok := nodesMap[x]; ok {
			if _, ok := nodesMap[y]; ok {
				break
			}
		}
	}

	return nodesMap[x].Level == nodesMap[y].Level && nodesMap[x].Parent != nodesMap[y].Parent
}
