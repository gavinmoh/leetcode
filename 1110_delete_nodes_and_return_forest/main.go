package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	forests := []*TreeNode{}

	deleteMap := make(map[int]bool)
	for _, val := range to_delete {
		deleteMap[val] = true
	}

	var dfs func(current, prev, currentRoot *TreeNode)
	dfs = func(current, prev, currentRoot *TreeNode) {
		if current == nil {
			return
		}

		if prev == nil && !deleteMap[currentRoot.Val] {
			forests = append(forests, currentRoot)
		}

		if deleteMap[current.Val] {
			if prev != nil {
				if prev.Left == current {
					prev.Left = nil
				} else {
					prev.Right = nil
				}
			}

			dfs(current.Left, nil, current.Left)
			dfs(current.Right, nil, current.Right)
		} else {
			dfs(current.Left, current, currentRoot)
			dfs(current.Right, current, currentRoot)
		}
	}

	dfs(root, nil, root)

	return forests
}
