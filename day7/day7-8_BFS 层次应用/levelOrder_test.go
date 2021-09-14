package day7_8

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) []int {
	res := make([]int,0,10)
	if root == nil {
		return res
	}

	//建立子树队列
	queue := make([]*TreeNode,0)
	queue = append(queue, root)

	//若子树队列中还有树则继续
	for len(queue) > 0{
		//建立存储当前层数据的list
		list := make([]int,0,len(queue)*2)
		//遍历当前队列中的每一个树，即遍历该层的所有树
		l := len(queue)
		for i:=0; i<l; i++{
			//取出队列第一个树
			level := queue[0]
			queue = queue[1:]	//记得弹出此树
			//将数据放入list
			list = append(list, level.Val)
			//判断左右子树并将其加入分支队列
			if level.Left != nil{
				queue = append(queue, level.Left)
			}
			if level.Right != nil{
				queue = append(queue, level.Right)
			}
		}
		//将该层数据送入res
		res = append(res,list...)
	}
	return res
}

func TestLevelOrder(t *testing.T){
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
	res := levelOrder(&root)
	assert.Equal(t, []int{3,9,20,15,7}, res, "Error in case 1")
}
