package day12_10_searchBST

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	for root != nil && root.Val != val{
		if root.Val > val{
			root = root.Left
			continue
		}
		root = root.Right
	}
	return root
}

func TestSearchBST(t *testing.T){
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
			Left:  nil,
			Right: nil,
		},
	}

	res := searchBST(&root, 2)
	assert.Equal(t, 2, res.Val, "Error in case 1")
}
