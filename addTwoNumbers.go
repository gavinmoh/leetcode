package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var newList *ListNode
	currentNode := &ListNode{}
	carryNumber := 0

	for l1 != nil || l2 != nil || carryNumber != 0 {
		var number1 int
		var number2 int

		if l1 != nil {
			number1 = l1.Val
			l1 = l1.Next
		} else {
			number1 = 0
		}

		if l2 != nil {
			number2 = l2.Val
			l2 = l2.Next
		} else {
			number2 = 0
		}

		sum := number1 + number2 + carryNumber

		if sum >= 10 {
			carryNumber = 1
			sum = sum - 10
		} else {
			carryNumber = 0 // reset carryNumber
		}

		currentNode.Val = sum
		currentNode.Next = &ListNode{}

		if newList == nil {
			newList = currentNode
		}

		if (l1 == nil && l2 == nil) && carryNumber == 0 {
			currentNode.Next = nil
			break
		}

		currentNode = currentNode.Next
	}
	return newList
}

func main() {
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}

	list := addTwoNumbers(l1, l2)
	for list != nil {
		fmt.Println(list.Val)
		list = list.Next
	}

}
