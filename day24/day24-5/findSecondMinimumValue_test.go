package day24_5

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findSecondMinimumValue(root *TreeNode) int {
	ans := -1
	rootVal := root.Val
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil || ans != -1 && node.Val >= ans {
			return
		}
		if node.Val > rootVal {
			ans = node.Val
		}
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return ans
}

func TestFindSecondMinimumValue(t *testing.T){
	//case 1
	root := TreeNode{
		Val:   2,
		Left:  &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   5,
			Left:  &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}
	res := findSecondMinimumValue(&root)
	assert.Equal(t, 5, res, "Error in case 2")
}
