package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(listOne, listTwo *ListNode) *ListNode {
	if listOne == nil {
		return listTwo
	}

	if listTwo == nil {
		return listOne
	}

	// if list 1 value is less than list 2 value
	// list 1 next is the result of merge list 1 next and list 2
	if listOne.Val < listTwo.Val {
		listOne.Next = MergeTwoLists(listOne.Next, listTwo)
		return listOne
	}

	// if list 2 value is less than list 1 value
	// list 2 next is the result of merge list 2 next and list 1
	listTwo.Next = MergeTwoLists(listOne, listTwo.Next)
	return listTwo
}

func MergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	var result *ListNode

	for _, list := range lists {
		result = MergeTwoLists(result, list)
	}

	return result

}
