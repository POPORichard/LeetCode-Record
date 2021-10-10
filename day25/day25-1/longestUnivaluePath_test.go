package day25_1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestUnivaluePath(root *TreeNode) int {
	res := 0
	var dfs func(*TreeNode, int, int)int

	dfs = func(root *TreeNode, f, path int) int {
		if root != nil{
			l := dfs(root.Left, root.Val, path)
			r := dfs(root.Right, root.Val, path)
			res = max(r+l, res)
			if f == root.Val{
				return max(l, r) + 1
			}
		}
		return 0
	}
	dfs(root, 0, 0)
	return res
}

func max (a,b int)int{
	if a>b {
		return a
	}
	return b
}

func TestLongestUnivaluePath(t *testing.T){
	//case 1
	root := TreeNode{
		Val:   1,
		Left:  &TreeNode{
			Val:   4,
			Left:  &TreeNode{
				Val:   4,
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
			Val:   5,
			Left:  nil,
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
	}
	res := longestUnivaluePath(&root)
	assert.Equal(t, 2, res, "Error in case 1")

}
