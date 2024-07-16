package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getDirections(root *TreeNode, startValue int, destValue int) string {
	var lca func(node *TreeNode, p int, q int) *TreeNode
	lca = func(node *TreeNode, p int, q int) *TreeNode {
		if node == nil {
			return nil
		}

		if node.Val == p || node.Val == q {
			return node
		}

		left := lca(node.Left, p, q)
		right := lca(node.Right, p, q)

		if left != nil && right != nil {
			return node
		}

		if left == nil {
			return right
		}

		return left
	}
	root = lca(root, startValue, destValue)

	var dfs func(node *TreeNode, target int, path []rune) []rune
	dfs = func(node *TreeNode, target int, path []rune) []rune {
		if node.Left != nil {
			leftPath := append(path, 'L')
			if node.Left.Val == target {
				return leftPath
			} else {
				result := dfs(node.Left, target, leftPath)
				if len(result) > 0 {
					return result
				}
			}
		}

		if node.Right != nil {
			rightPath := append(path, 'R')
			if node.Right.Val == target {
				return rightPath
			} else {
				result := dfs(node.Right, target, rightPath)
				if len(result) > 0 {
					return result
				}
			}
		}

		return []rune{}
	}

	sPath := dfs(root, startValue, []rune{})
	dPath := dfs(root, destValue, []rune{})

	for i := range sPath {
		sPath[i] = 'U'
	}

	return string(sPath) + string(dPath)
}
