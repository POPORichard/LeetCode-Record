package day22_9

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

func getMinimumDifference(root *TreeNode) int {
	ans, pre := math.MaxInt64, -1
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if pre != -1 && node.Val-pre < ans {
			ans = node.Val - pre
		}
		pre = node.Val
		dfs(node.Right)
	}
	dfs(root)
	return ans
}

func TestGetMinimumDifference(t *testing.T){
	root := TreeNode{
		Val:   1,
		Left:  nil,
		Right: &TreeNode{
			Val:   3,
			Left:  &TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}
	res := getMinimumDifference(&root)
	assert.Equal(t, 1, res, "Error in case 1")

}
