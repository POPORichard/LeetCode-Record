package day9_1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {
	//建立队列
	stack := make([]*TreeNode,0, 10)
	stack = append(stack, root)
	num := 1

	//当队列不为空时
	for len(stack) >0{
		//l存储下一层有多少节点
		l := 0
		//sum存储本层的和
		sum := 0
		for i:=0; i<num; i++{

			//出队列
			root = stack[0]
			stack = stack[1:]

			if root.Left != nil{
				stack = append(stack, root.Left)
				l++
			}
			if root.Right != nil{
				stack = append(stack, root.Right)
				l++
			}

			sum += root.Val
		}
		//如果下一层没有节点了，则返回
		if l == 0 {
			return sum
		}
		//将下一层的节点数给num
		num = l
	}
	return -1
}

func TestDeepestLeavesSum(t *testing.T){
	//case 1
	root := TreeNode{
		Val:   1,
		Left:  &TreeNode{
			Val:   2,
			Left:  &TreeNode{
				Val:   4,
				Left:  &TreeNode{
					Val:   7,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: &TreeNode{
				Val:   6,
				Left:  nil,
				Right: &TreeNode{
					Val:   8,
					Left:  nil,
					Right: nil,
				},
			},
		},
	}
	res := deepestLeavesSum(&root)
	assert.Equal(t, 15, res, "Error in case 1")

}
