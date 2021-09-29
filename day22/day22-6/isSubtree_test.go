package day22_6

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil{
		return false
	}
	return sameTreeNode(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func sameTreeNode(root, sub *TreeNode) bool{
	if  root == nil && sub == nil{
		return true
	}
	if root == nil || sub == nil{
		return false
	}
	if root.Val != sub.Val {
		return false
	}else{
		return sameTreeNode(root.Left, sub.Left) && sameTreeNode(root.Right, sub.Right)
	}
}

func TestIsSubtree(t *testing.T){
	root := TreeNode{
		Val:   3,
		Left:  &TreeNode{
			Val:   4,
			Left:  &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   5,
			Left:  nil,
			Right: nil,
		},
	}
	subRoot := TreeNode{
		Val:   4,
		Left:  &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	}
	res := isSubtree(&root, &subRoot)
	assert.Equal(t, true, res, "Error in case 1")
}