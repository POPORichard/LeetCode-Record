package day9_6

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}

	//设置初始头节点
	start := head

	//进入循环
	for {
		//记录上一个节点为last
		last := head
		//头节点后移
		head = head.Next

		//如果头节点小于上一个节点，则需要排序
		if head.Val < last.Val {
			//先将last节点指向head节点的下一个节点以从链表中移出head
			last.Next = head.Next
			//判断当前节点的值是否比第一个节点还小
			if head.Val < start.Val {
				//将head作为新的开始
				head.Next = start
				start = head
			} else {
				//建立finder节点及finder节点的上一个节点finderLast
				finderLast := start
				finder := start.Next
				//从头开始寻找一个大于head的节点
				for {
					if head.Val < finder.Val {
						//若找到，将head重新连接进入
						finderLast.Next = head
						head.Next = finder
						break
					}
					//finder后移
					finderLast = finder
					finder = finder.Next
				}
			}
			//回到弹出head的位置
			head = last
		}
		//结束条件
		if head.Next == nil {
			break
		}
	}
	return start
}

func TestInsertionSortList(t *testing.T) {
	//case 1
	head := ListNode{
		Val: -1,
		Next: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  0,
						Next: nil,
					},
				},
			},
		},
	}
	head = *insertionSortList(&head)
	res := make([]int, 0, 10)
	for {
		res = append(res, head.Val)
		if head.Next == nil {
			break
		}
		head = *head.Next
	}
	assert.Equal(t, []int{-1, 0, 3, 4, 5}, res, "Error in case 1")

}
