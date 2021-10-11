package day26_5

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func numColor(root *TreeNode) int {
	hashMap := map[int]struct{}{}

	var dfs func(*TreeNode)

	dfs = func(node *TreeNode) {
		if _,ok := hashMap[node.Val]; !ok{
			hashMap[node.Val] = struct{}{}
		}
		if node.Right != nil{
			dfs(node.Right)
		}
		if node.Left != nil{
			dfs(node.Left)
		}
	}

	dfs(root)
	return len(hashMap)
}

func TestNumColor(t *testing.T){
	//case 1
	root := TreeNode{
		Val:   1,
		Left:  &TreeNode{
			Val:   3,
			Left:  &TreeNode{
				Val:   1,
				Left:  &TreeNode{
					Val:   2,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
			Right: nil,
		},
		Right: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	}
	res := numColor(&root)
	assert.Equal(t, 3, res, "Error in case 1")
}
