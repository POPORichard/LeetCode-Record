package day9_4

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
//逐层遍历，找到第一个没有子节点的节点，返回深度
func minDepth(root *TreeNode) int {
	if root == nil {return 0}

	//建立队列
	queue := make([]*TreeNode,0,10)
	queue = append(queue, root)
	num := 1
	depth := 1

	for queue != nil{
		l := 0
		for i:=0; i<num; i++{
			//flag判断当前节点是否具有子节点
			flag := false
			root = queue[0]
			queue = queue[1:]

			if root.Left != nil {
				queue = append(queue, root.Left)
				l++
				flag = true
			}
			if root.Right != nil {
				queue = append(queue, root.Right)
				l++
				flag = true
			}

			//若当前节点没有子节点，则该节点为最浅节点
			if !flag {
				return depth
			}
		}
		depth++
		num = l
	}
	return depth
}

func TestMinDepth(t *testing.T){
	//case 1
	root := TreeNode{
		Val:   3,
		Left:  &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{
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
	res :=minDepth(&root)
	assert.Equal(t, 2, res, "Error in case 1")
}
