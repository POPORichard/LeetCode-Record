package day9_10

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	target := head
	for i:=1; i<k; i++{
		head = head.Next
	}
	for {
		if head.Next == nil{
			return target
		}
		head = head.Next
		target = target.Next
	}
}

func TestGetKthFromEnd(t *testing.T){
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
	res := getKthFromEnd(&head, 2)
	num := make([]int,0,5)
	for{
		num = append(num, res.Val)
		if res.Next == nil {break}
		res = res.Next
	}
	assert.Equal(t, []int{4,5}, num,"Error in case 1")

}
