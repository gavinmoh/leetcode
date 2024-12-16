package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func reverseOddLevels(root *TreeNode) *TreeNode {
	tree := map[int][]*TreeNode{}
	var dfs func(current *TreeNode, level int)
	dfs = func(current *TreeNode, level int) {
		if current == nil {
			return
		}

		if _, ok := tree[level]; !ok {
			tree[level] = []*TreeNode{}
		}

		tree[level] = append(tree[level], current)

		dfs(current.Left, level+1)
		dfs(current.Right, level+1)
	}

	dfs(root, 0)

	for level, nodes := range tree {
		if level%2 == 0 {
			continue
		}

		values := []int{}
		for _, node := range nodes {
			values = append(values, node.Val)
		}

		i := len(values) - 1
		for _, node := range nodes {
			node.Val = values[i]
			i--
		}
	}

	return root
}
