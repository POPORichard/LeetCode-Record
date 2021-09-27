package day20_5_getIntersectionNode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
//设a在合并前的长度为a
//设b在合并前的长度为b
//设公共的长度为c
//则，pa指针跑a+c+b时与pb指针跑b+c+a时，相遇再交汇处
func betterGetIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pa, pb := headA, headB
	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}
		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}
	return pa
}

func TestBetterGetIntersectionNode(t *testing.T){
	//case 1
	intersect := ListNode{
		Val:  8,
		Next: &ListNode{
			Val:  4,
			Next: &ListNode{
				Val:  5,
				Next: nil,
			},
		},
	}
	headA := ListNode{
		Val:  4,
		Next: &ListNode{
			Val:  1,
			Next: &intersect,
		},
	}
	headB := ListNode{
		Val:  5,
		Next: &ListNode{
			Val:  0,
			Next: &ListNode{
				Val:  1,
				Next: &intersect,
			},
		},
	}
	res := betterGetIntersectionNode(&headA, &headB)
	assert.Equal(t, 8, res.Val, "Error in case 1")
}