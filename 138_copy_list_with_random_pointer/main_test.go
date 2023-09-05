package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCopyRandomList(t *testing.T) {
	// create the list
	node1 := &Node{Val: 7}
	node2 := &Node{Val: 13}
	node3 := &Node{Val: 11}
	node4 := &Node{Val: 10}
	node5 := &Node{Val: 1}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5

	node2.Random = node1
	node3.Random = node5
	node4.Random = node3
	node5.Random = node1

	// copy the list
	newHead := CopyRandomList(node1)
	newNode1 := newHead
	newNode2 := newHead.Next
	newNode3 := newHead.Next.Next
	newNode4 := newHead.Next.Next.Next
	newNode5 := newHead.Next.Next.Next.Next

	// check the list
	assert.Equal(t, 7, newNode1.Val)
	assert.Equal(t, 13, newNode2.Val)
	assert.Equal(t, 11, newNode3.Val)
	assert.Equal(t, 10, newNode4.Val)
	assert.Equal(t, 1, newNode5.Val)

	assert.Equal(t, newNode1, newNode2.Random)
	assert.Equal(t, newNode5, newNode3.Random)
	assert.Equal(t, newNode3, newNode4.Random)
	assert.Equal(t, newNode1, newNode5.Random)

}
