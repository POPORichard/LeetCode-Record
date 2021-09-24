package day17_3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {return nil}
	res := head
	for head.Next != nil{
		if head.Next.Val != head.Val{
			head = head.Next
			continue
		}
		head.Next = head.Next.Next
	}
	return res
}

func TestDeleteDuplicates(t *testing.T){
	//case 1
	head := ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  1,
			Next: &ListNode{
				Val:  2,
				Next: &ListNode{
					Val:  3,
					Next: &ListNode{
						Val:  3,
						Next: nil,
					},
				},
			},
		},
	}
	res := deleteDuplicates(&head)
	vals := make([]int,0,10)
	for res != nil{
		vals = append(vals, res.Val)
		res = res.Next
	}
	assert.Equal(t, []int{1,2,3}, vals, "Error in case 1")
}
