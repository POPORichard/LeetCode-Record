package day20

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

func TestDeleteNode(t *testing.T){

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
	res := []int{}
	for {
		res = append(res, head.Val)
		if head.Next == nil{
			break
		}
		head = *head.Next
	}
	assert.Equal(t, []int{4,1,9}, res, "Error in case 1")
}
