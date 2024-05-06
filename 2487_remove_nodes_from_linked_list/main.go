package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type Stack []*ListNode

func (s *Stack) Size() int {
	return len(*s)
}

func (s *Stack) Top() *ListNode {
	arr, n := *s, s.Size()
	return arr[n-1]
}

func (s *Stack) Pop() *ListNode {
	arr, n := *s, s.Size()
	last := arr[n-1]
	(*s) = arr[:n-1]

	return last
}

func (s *Stack) Push(node *ListNode) {
	(*s) = append((*s), node)
}

func removeNodes(head *ListNode) *ListNode {
	stack := &Stack{}
	current := head
	for current != nil {
		for stack.Size() > 0 && stack.Top().Val < current.Val {
			stack.Pop()
		}
		stack.Push(current)
		current = current.Next
	}

	// reconstructing the list
	head = stack.Pop()
	head.Next = nil
	for stack.Size() > 0 {
		prev := head
		head = stack.Pop()
		head.Next = prev
	}

	return head
}
