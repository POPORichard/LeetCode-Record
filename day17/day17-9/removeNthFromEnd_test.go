package day17_9

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	start := head
	toucher := head

	for i:=0; i<n; i++{
		toucher = toucher.Next
	}
	if toucher == nil{
		return head.Next
	}



	for {
		if toucher.Next == nil{
			head.Next = head.Next.Next
			return start
		}
		head = head.Next
		toucher = toucher.Next

	}
}

func TestRemoveNthFromEnd(t *testing.T){
	//case 1
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
	res := removeNthFromEnd(&head, 4)
	ans := make([]int, 0, 10)
	for res !=nil{
		ans = append(ans, res.Val)
		res = res.Next
	}
	assert.Equal(t, []int{1,3,4,5}, ans, "Error in case 1")
}
