package day21_10

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumRootToLeaf(root *TreeNode) int {
	sum := 0
	dfs(root, 0, &sum)
	return sum
}
func dfs(root *TreeNode, cont int, sum *int) {
	if root == nil {
		return
	}
	cont = (cont << 1) + root.Val
	if root.Left == nil && root.Right == nil {
		*sum = *sum + cont
	}
	dfs(root.Left, cont, sum)
	dfs(root.Right, cont, sum)

}

func TestSumRootToLeaf(t *testing.T){
	//case 1
	root := TreeNode{
		Val:   1,
		Left:  &TreeNode{
			Val:   0,
			Left:  &TreeNode{
				Val:   0,
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
			Left:  &TreeNode{
				Val:   0,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
		},
	}

	res :=sumRootToLeaf(&root)
	assert.Equal(t, 22, res, "Error in case 1")
}
