package day10_3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	nums := make([]int, 0, 10)
	for{
		nums = append(nums,head.Val)
		if head.Next == nil{
			break
		}
		head = head.Next
	}
	left := 0
	right := len(nums)-1
	for{
		if left >= right{return true}
		if nums[left] != nums[right]{return false}
		left++
		right--
	}
}

func TestIsPalindrome(t *testing.T){
	//case 1
	head := ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{
				Val:  2,
				Next: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
		},
	}
	res := isPalindrome(&head)
	assert.Equal(t, true, res, "Error in case 1")
}
