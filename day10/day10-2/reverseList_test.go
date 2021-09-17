package day10_2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {return nil}
	var prev *ListNode
	next := head.Next
	for {
		head.Next = prev
		prev = head
		if next == nil{break}
		head = next
		next = head.Next
	}
	return head
}

func TestReverseList(t *testing.T){
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
	res := reverseList(&head)
	nums := make([]int,0,10)
	for {
		if res == nil{break}
		nums = append(nums, res.Val)
		res = res.Next
	}
	assert.Equal(t,[]int{5,4,3,2,1}, nums, "Error in case 1")
}
