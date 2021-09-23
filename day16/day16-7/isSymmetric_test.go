package day16_7

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {return true}
	return check(root.Right, root.Left)
}

func check(right, left *TreeNode) bool{
	if right == nil && left == nil{
		return true
	}
	if right != nil && left != nil{
		flag := true
		if right.Val != left.Val{
			return false
		}

		flag = check(right.Right, left.Left)
		if flag == false{
			return false
		}

		flag = check(right.Left, left.Right)
		if flag == false{
			return false
		}

		return true
	}

	return false
}

func TestIsSymmetric(t *testing.T){
	root := TreeNode{
		Val:   1,
		Left:  &TreeNode{
			Val:   2,
			Left:  &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   2,
			Left:  &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
	}
	res := isSymmetric(&root)
	assert.Equal(t, true, res, "Error in case 1")

}


