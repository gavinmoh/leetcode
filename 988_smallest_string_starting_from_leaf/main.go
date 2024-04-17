package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func smallestFromLeaf(root *TreeNode) string {
	result := []byte{}
	var dfs func(node *TreeNode, acc []byte)
	dfs = func(node *TreeNode, acc []byte) {
		newAcc := append([]byte{byte('a' + node.Val)}, acc...)
		if node.Left == nil && node.Right == nil {
			if len(result) == 0 {
				result = newAcc
			} else {
				result = min(result, newAcc)
			}
		}
		if node.Left != nil {
			dfs(node.Left, newAcc)
		}
		if node.Right != nil {
			dfs(node.Right, newAcc)
		}
	}
	dfs(root, []byte{})

	return string(result)
}

func min(a, b []byte) []byte {
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] == b[i] {
			continue
		}

		if a[i] < b[i] {
			return a
		}

		return b
	}

	if len(a) < len(b) {
		return a
	}

	return b
}
