package day5_1

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var paths []string

func binaryTreePaths(root *TreeNode) []string {
	paths = []string{}
	constructPaths(root, "")
	return paths
}

func constructPaths(root *TreeNode, path string) {
	if root != nil {
		pathSB := path
		pathSB += strconv.Itoa(root.Val)
		if root.Left == nil && root.Right == nil {
			paths = append(paths, pathSB)
		} else {
			pathSB += "->"
			constructPaths(root.Left, pathSB)
			constructPaths(root.Right, pathSB)
		}
	}
}

func TestBinaryTreePaths(t *testing.T){
	//case 1
	input := TreeNode{
		Val:   1,
		Left:  &TreeNode{
			Val:   2,
			Left:  nil,
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}
	res := binaryTreePaths(&input)
	assert.Equal(t, []string{"1->2->5","1->3"}, res, "Error in case 1")

}
