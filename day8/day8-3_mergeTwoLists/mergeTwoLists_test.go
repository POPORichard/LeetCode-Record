package day8_3_mergeTwoLists

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	//建立哑巴节点
	dummy := &ListNode{Val: 0}
	//头指针此时指向哑巴节点
	head := dummy
	//比较大小并添加链接
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			head.Next = l1
			l1 = l1.Next
		} else {
			head.Next = l2
			l2 = l2.Next
		}
		//将头指针后移
		head = head.Next
	}
	// 连接l1 未处理完节点
	for l1 != nil {
		head.Next = l1
		head = head.Next
		l1 = l1.Next
	}
	// 连接l2 未处理完节点
	for l2 != nil {
		head.Next = l2
		head = head.Next
		l2 = l2.Next
	}
	return dummy.Next
}

func TestMergeTwoLists(t *testing.T){
	l1 := ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	l2 := ListNode{
		Val:  1,
		Next: &ListNode{
			Val:  3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	node := mergeTwoLists(&l1,&l2)
	res := make([]int,0,10)
	for node != nil{
		res = append(res, node.Val)
		node = node.Next
	}
	assert.Equal(t, []int{1,1,2,3,4,4}, res, "Error in case 1")

}
