package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	var dfs func(preLeft, preRight, postLeft int) *TreeNode
	dfs = func(preLeft, preRight, postLeft int) *TreeNode {
		if preLeft > preRight {
			return nil
		}

		if preLeft == preRight {
			return &TreeNode{Val: preorder[preLeft]}
		}

		leftRootVal := preorder[preLeft+1]
		leftNodesCount := 1
		for postorder[postLeft+leftNodesCount-1] != leftRootVal {
			leftNodesCount++
		}

		return &TreeNode{
			Val:   preorder[preLeft],
			Left:  dfs(preLeft+1, preLeft+leftNodesCount, postLeft),
			Right: dfs(preLeft+leftNodesCount+1, preRight, postLeft+leftNodesCount),
		}
	}

	return dfs(0, len(preorder)-1, 0)
}
