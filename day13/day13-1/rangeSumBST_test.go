package day13_1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	sum := 0
	//创建队列
	q := []*TreeNode{root}
	for len(q) > 0 {
		//队列出
		node := q[0]
		q = q[1:]

		if node == nil {
			continue
		}

		if node.Val > high {
			q = append(q, node.Left)
		} else if node.Val < low {
			q = append(q, node.Right)
		} else {
			sum += node.Val
			q = append(q, node.Left, node.Right)
		}
	}
	return sum
}

func TestRangeSumBST(t *testing.T){
	//case 1
	root := TreeNode{
		Val:   10,
		Left:  &TreeNode{
			Val:   5,
			Left:  &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   15,
			Left:  nil,
			Right: &TreeNode{
				Val:   18,
				Left:  nil,
				Right: nil,
			},
		},
	}
	res := rangeSumBST(&root,7,15)
	assert.Equal(t, 32, res, "Error in case 1")

}
