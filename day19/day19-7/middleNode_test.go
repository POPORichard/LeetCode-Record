package day19_7

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {

	mid := head

	for {
		if head.Next == nil {
			break
		}
		head = head.Next
		if head.Next == nil {
			mid = mid.Next
			break
		}
		head = head.Next

		mid = mid.Next

	}

	return mid

}

func TestMiddleNode(t *testing.T){
	head := ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{
				Val:  3,
				Next: &ListNode{
					Val:  4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	res := middleNode(&head)
	assert.Equal(t, 3, res.Val, "Error in case 1")
}
