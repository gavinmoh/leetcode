package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func amountOfTime(root *TreeNode, start int) int {
	graph := make(map[int][]int)

	// converting tree to graph
	var treeToGraph func(*TreeNode, int)
	treeToGraph = func(node *TreeNode, parent int) {
		if node == nil {
			return
		}

		if _, ok := graph[node.Val]; !ok {
			graph[node.Val] = []int{}
		}

		if parent != -1 {
			graph[node.Val] = append(graph[node.Val], parent)
		}

		if node.Left != nil {
			graph[node.Val] = append(graph[node.Val], node.Left.Val)
		}

		if node.Right != nil {
			graph[node.Val] = append(graph[node.Val], node.Right.Val)
		}
		treeToGraph(node.Left, node.Val)
		treeToGraph(node.Right, node.Val)
	}
	treeToGraph(root, -1)

	queue := []int{start}
	visited := make(map[int]bool)
	depth := 0

	visited[start] = true
	for len(queue) > 0 {
		newQueue := []int{}
		for _, num := range queue {
			adjacentNodes := graph[num]
			for _, node := range adjacentNodes {
				if _, ok := visited[node]; !ok {
					visited[node] = true
					newQueue = append(newQueue, node)
				}
			}
		}
		depth += 1
		queue = newQueue
	}

	return depth - 1
}
