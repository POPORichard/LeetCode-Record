package day17_4

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
	if head.Next == nil {return head}
	if head == nil {
		return nil
	}

	dummy := &ListNode{0, head}

	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			x := cur.Next.Val
			for cur.Next != nil && cur.Next.Val == x {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}

	return dummy.Next


}

func TestDeleteDuplicates(t *testing.T){
	//case 1
	head := ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  1,
			Next: &ListNode{
				Val:  1,
				Next: &ListNode{
					Val:  2,
					Next: &ListNode{
						Val:  3,
						Next: nil,
					},
				},
			},
		},
	}
	res := deleteDuplicates(&head)
	val := make([]int,0,10)
	for res != nil{
		val = append(val, res.Val)
		res =res.Next
	}
	assert.Equal(t, []int{2,3}, val, "Error in case 1")

	//case 2
	head = ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{
				Val:  3,
				Next: &ListNode{
					Val:  3,
					Next: &ListNode{
						Val:  4,
						Next: &ListNode{
							Val:  4,
							Next: &ListNode{
								Val:  5,
								Next: nil,
							},
						},
					},
				},
			},
		},
	}
	res = deleteDuplicates(&head)
	val = make([]int,0,10)
	for res != nil{
		val = append(val, res.Val)
		res =res.Next
	}
	assert.Equal(t, []int{1,2,5}, val, "Error in case 2")
}

