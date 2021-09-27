package day20_5_getIntersectionNode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	hashMap :=map[*ListNode]struct{}{}

	for headA != nil{
		hashMap[headA] = struct {}{}
			headA = headA.Next
	}

	for headB != nil{
		if _,ok := hashMap[headB]; ok{
			return headB
		}
		headB = headB.Next
	}
	return nil
}

func TestGetIntersectionNode(t *testing.T){
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
	res := getIntersectionNode(&headA, &headB)
	assert.Equal(t, 8, res.Val, "Error in case 1")
}


