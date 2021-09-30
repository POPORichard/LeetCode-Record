package day23_8

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left >= right || left < 1 || right < 1 {
		return head
	}
	dummyNode := ListNode{Val: -1}
	leftHead := &dummyNode
	curNode := head
	dummyNode.Next = curNode
	rightHead := curNode
	for i := 1; i < left && curNode != nil; i++ {
		leftHead = curNode
		curNode = curNode.Next
		rightHead = curNode
	}
	for i := -1; i < right-left && curNode != nil; i++ {
		tmp := curNode
		curNode = curNode.Next
		tmp.Next = leftHead.Next
		leftHead.Next = tmp
	}
	rightHead.Next = curNode
	return dummyNode.Next
}

func TestReverseBetween(t *testing.T){
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
	res := reverseBetween(&head, 2,4)
	ans := make([]int,0,10)
	for res != nil{
		ans = append(ans, res.Val)
		res = res.Next
	}
	assert.Equal(t, []int{1,4,3,2,5}, ans,"Error in case 1")
}
