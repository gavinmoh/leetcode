package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type FindElements struct {
	Root  *TreeNode
	Nodes map[int]bool
}

func Constructor(root *TreeNode) FindElements {
	root.Val = 0
	obj := FindElements{
		Root:  root,
		Nodes: map[int]bool{0: true},
	}
	obj.Recover()

	return obj
}

func (this *FindElements) Find(target int) bool {
	return this.Nodes[target]
}

func (this *FindElements) Recover() {
	this.dfs(this.Root)
}

func (this *FindElements) dfs(node *TreeNode) {
	if node == nil {
		return
	}

	if node.Left != nil && node.Left.Val == -1 {
		node.Left.Val = 2*node.Val + 1
		this.Nodes[node.Left.Val] = true
	}
	this.dfs(node.Left)

	if node.Right != nil && node.Right.Val == -1 {
		node.Right.Val = 2*node.Val + 2
		this.Nodes[node.Right.Val] = true
	}
	this.dfs(node.Right)
}

/**
* Your FindElements object will be instantiated and called as such:
* obj := Constructor(root);
* param_1 := obj.Find(target);
 */
