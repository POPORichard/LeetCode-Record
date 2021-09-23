package day11_5_detectCycle

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {return head}
	faster := head
	slower := head

	for {
		if faster.Next == nil{return nil}
		faster = faster.Next
		if faster.Next == nil{return nil}
		faster = faster.Next
		slower = slower.Next
		if slower == faster {
			break
		}
	}
	for {
		if head == slower{
			return head
		}
		head = head.Next
		slower = slower.Next
	}
}

func TestDetectCycle(t *testing.T){
//case 1
	head := ListNode{
		Val:  3,
		Next: nil,
	}
	loop := ListNode{
		Val:  2,
		Next: &ListNode{
			Val:  0,
			Next: &ListNode{
				Val:  -4,
				Next: nil,
			},
		},
	}
	loop.Next.Next = &loop
	head.Next = &loop

	res := detectCycle(&head)
	assert.Equal(t, 2, res.Val, "Error in case 1")

	//case 2
	head = ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}
	head.Next.Next = &head
	res = detectCycle(&head)
	assert.Equal(t, 1, res.Val, "Error in case 2")

}
