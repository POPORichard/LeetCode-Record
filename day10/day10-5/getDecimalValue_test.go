package day10_5

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	ans:=0
	root:=head
	for root!=nil {
		ans=ans<<1+root.Val
		root=root.Next
	}
	return ans
}

func TestGetDecimalValue(t *testing.T){
	//case 1
	head := ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  0,
			Next: &ListNode{
				Val:  1,
				Next: nil,
			},
		},
	}
	res := getDecimalValue(&head)
	assert.Equal(t, 5, res, "Error in case 1")

}
