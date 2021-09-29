package day22_7

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxLevelSum(root *TreeNode) int {
	if root == nil {return 0}
	sum := make([]int,0)

	var levelSum func(*TreeNode, int)
	levelSum = func(r *TreeNode, level int) {
		if r == nil{
			return
		}

		if len(sum)-1 < level{
			sum= append(sum,0)
		}

		sum[level] += r.Val

		levelSum(r.Left, level+1)
		levelSum(r.Right, level+1)
	}

	levelSum(root, 0)
	fmt.Println(sum)

	if len(sum) == 1{return sum[0]}

	res := len(sum)-1
	biggest := sum[len(sum)-1]
	for i:= len(sum)-2; i>=0; i--{
		if sum[i] >= biggest{
			res = i
			biggest = sum[i]
		}
	}
	return res+1
}



//func levelSum1(root *TreeNode, sum *[]int, level int){
//	arr := *sum
//	if root == nil{
//		return
//	}
//
//	if len(arr)-1 < level{
//		arr= append(arr,0)
//	}
//
//	arr[level] += root.Val
//
//	levelSum(root.Left, &arr, level+1)
//	levelSum(root.Right, &arr, level+1)
//}

func TestMaxLevelSum(t *testing.T){
	root := TreeNode{
		Val:   1,
		Left:  &TreeNode{
			Val:   7,
			Left:  &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   -8,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   0,
			Left:  nil,
			Right: nil,
		},
	}
	res := maxLevelSum(&root)
	assert.Equal(t, 2, res, "Error in case 1")

	//case 1
	root = TreeNode{
		Val:   1,
		Left:  &TreeNode{
			Val:   1,
			Left:  &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   -8,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   0,
			Left:  &TreeNode{
				Val:   -7,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   9,
				Left:  nil,
				Right: nil,
			},
		},
	}
	res = maxLevelSum(&root)
	assert.Equal(t, 1, res, "Error in case 1")
}
