package day20_9

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return validBST(root, math.MinInt64, math.MaxInt64)
}

func validBST(root *TreeNode, lower, upper int) bool {
	if root == nil {return true}

	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= upper {
		return false
	}
	return validBST(root.Left, lower, root.Val) && validBST(root.Right, root.Val, upper)
}


func TestIsValidBST(t *testing.T){
	//case 1
	root := TreeNode{
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
	}
	res := isValidBST(&root)
	assert.Equal(t, res, true, "Error in case 1")

	//case 2
	root = TreeNode{
		Val:   0,
		Left:  &TreeNode{
			Val:   -1,
			Left:  nil,
			Right: nil,
		},
		Right: nil,
	}
	res = isValidBST(&root)
	assert.Equal(t, res, true, "Error in case 2")
}
