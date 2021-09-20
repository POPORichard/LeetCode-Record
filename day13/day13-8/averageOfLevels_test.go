package day13_8

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//counts 用于存储二叉树的每一层的节点数
//sums 用于存储二叉树的每一层的节点值之和
type data struct{ sum, count int }

func averageOfLevels(root *TreeNode) []float64 {
	levelData := []data{}

	//建立深度搜索函数
	var dfs func(node *TreeNode, level int)
	dfs = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		if level < len(levelData) {
			levelData[level].sum += node.Val
			levelData[level].count++
		} else {
			levelData = append(levelData, data{node.Val, 1})
		}
		//深度递归
		dfs(node.Left, level+1)
		dfs(node.Right, level+1)
	}

	dfs(root, 0)

	//计算平均值
	averages := make([]float64, len(levelData))
	for i, d := range levelData {
		averages[i] = float64(d.sum) / float64(d.count)
	}

	return averages
}

func TestAverageOfLevels(t *testing.T){
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
	res := averageOfLevels(&root)
	assert.Equal(t, []float64{3, 14.5, 11}, res, "Error in case 1")
}
