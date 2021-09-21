package day14_1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left != nil && root.Val != root.Left.Val  || root.Right != nil && root.Val != root.Right.Val {
		return false
	}
	return isUnivalTree(root.Left) && isUnivalTree(root.Right)
}

func TestIsUnivalTree(t *testing.T){
	//case 1
	root := TreeNode{
		Val:   1,
		Left:  &TreeNode{
			Val:   1,
			Left:  &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
		},
	}
	res := isUnivalTree(&root)
	assert.Equal(t, true, res ,"Error in case 1")
}
