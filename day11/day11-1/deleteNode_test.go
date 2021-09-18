package day11_1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	next := node.Next
	node.Next = next.Next
	node.Val = next.Val
}

func TestDeleteNode(t *testing.T){
	//case 1
	node := ListNode{
		Val:  5,
		Next: &ListNode{
			Val:  1,
			Next: &ListNode{
				Val:  9,
				Next: nil,
			},
		},
	}

	head := ListNode{
		Val:  4,
		Next: &node,
	}

	deleteNode(&node)
	res := make([]int,0,10)
	for {
		res = append(res, head.Val)
		if head.Next == nil{break}
		head = *head.Next
	}
	assert.Equal(t, []int{4,1,9}, res, "Error in case 1")
}
