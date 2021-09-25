package day18_1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := invertTree(root.Left)
	right := invertTree(root.Right)
	root.Left = right
	root.Right = left
	return root
}

func TestInvertTree(t *testing.T){
	//case 1
	root := TreeNode{
		Val:   4,
		Left:  &TreeNode{
			Val:   2,
			Left:  &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   7,
			Left:  &TreeNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   9,
				Left:  nil,
				Right: nil,
			},
		},
	}
	res := invertTree(&root)
	assert.Equal(t, 4, res.Val, "Error in case 1")
}