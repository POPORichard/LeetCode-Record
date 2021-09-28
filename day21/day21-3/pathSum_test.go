package day21_3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
//前缀和

func pathSum(root *TreeNode, targetSum int)(ans int) {

	preSum := map[int]int{0: 1}

	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, curr int) {
		if node == nil {
			return
		}
		curr += node.Val
		ans += preSum[curr-targetSum]
		preSum[curr]++
		dfs(node.Left, curr)
		dfs(node.Right, curr)
		preSum[curr]--
		return
	}

	dfs(root, 0)

	return

}

func TestPathSum(t *testing.T){
	root := TreeNode{
		Val:   10,
		Left:  &TreeNode{
			Val:   5,
			Left:  &TreeNode{
				Val:   3,
				Left:  &TreeNode{
					Val:   3,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   -2,
					Left:  nil,
					Right: nil,
				},
			},
			Right: &TreeNode{
				Val:   2,
				Left:  nil,
				Right: &TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
			},
		},
		Right: &TreeNode{
			Val:   -3,
			Left:  nil,
			Right: &TreeNode{
				Val:   11,
				Left:  nil,
				Right: nil,
			},
		},
	}
	res := pathSum(&root, 8)
	assert.Equal(t, 3, res, "Error in case 1")
}
