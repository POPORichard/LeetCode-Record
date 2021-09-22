package day15_1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
type ListNode struct {
	Val  int
	Next *ListNode
}

func splitListToParts(head *ListNode, k int) []*ListNode {
	length := 0
	pointer := head
	for pointer != nil{
		pointer = pointer.Next
		length++
	}

	partLength := length/k
	longerPart := length%k

	ans := make([]*ListNode,k)
	pointer = head
	t:=0

	for i:=0; i<k && pointer != nil; i++{
		ans[i] = pointer

		if i<longerPart{
			t=-1
		}else {
			t=0
		}
		for ; t<partLength-1; t++{
			pointer = pointer.Next
		}
		pointer, pointer.Next = pointer.Next, nil
	}
	return ans
}

func TestSplitListToParts(t *testing.T){
	//case 1
	head := ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	res := splitListToParts(&head, 5)

	assert.Equal(t, 1, res[0].Val, "Error in case 1")
}