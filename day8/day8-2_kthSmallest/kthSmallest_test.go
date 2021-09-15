package day8_2_kthSmallest

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	stack := make([]*TreeNode,0,10)		//建立栈
	for len(stack)>0 || root !=nil{
		for root != nil{
			stack = append(stack, root)	//将root入栈
			root = root.Left			//一直左走
		}

		tree := stack[len(stack)-1]		//出栈操作
		stack = stack[:len(stack)-1]

		if tree != nil{
			k--
			if k ==0{
				return tree.Val
			}
		}
		root = tree.Right	//进入右树
	}
	return -1
}

func TestKthSmallest(t *testing.T){

	//case 1
	root := TreeNode{
		Val:   5,
		Left:  &TreeNode{
			Val:   3,
			Left:  &TreeNode{
				Val:   2,
				Left:  &TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   6,
			Left:  nil,
			Right: nil,
		},
	}

	res := kthSmallest(&root,3)
	assert.Equal(t, 3, res, "Error in case 1")
}
