package day8_10

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	//建立队列
	stack := make([]*TreeNode, 0, 10)
	stack = append(stack, root)
	ans := make([][]int, 0, 10)
	num := 1
	for len(stack) > 0 && root != nil {
		list := make([]int, 0, 10)
		//记录每层节点数
		l := 0
		for i := 0; i < num; i++ {
			//队列出
			root = stack[0]
			stack = stack[1:]

			list = append(list, root.Val)

			//队列入
			if root.Left != nil {
				stack = append(stack, root.Left)
				l++
			}
			if root.Right != nil {
				stack = append(stack, root.Right)
				l++
			}
		}
		num = l
		ans = append(ans, list)
	}
	return ans
}

func TestLevelOrder(t *testing.T) {
	//case 1
	root := TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val:   15,
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

	res := levelOrder(&root)
	assert.Equal(t, [][]int{{3}, {9, 20}, {15, 7}}, res, "Error in case 1")
}
