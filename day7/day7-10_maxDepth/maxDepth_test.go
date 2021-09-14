package day7_10_maxDepth

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	if root.Children == nil {
		return 1
	}

	list := make([]int, 0, len(root.Children))
	for i := range root.Children {
		list = append(list, 1+maxDepth(root.Children[i]))
	}

	if len(list) == 1 {
		return list[0]
	}

	biggest := 0
	for i := range list {
		if list[i] > biggest {
			biggest = list[i]
		}
	}
	return biggest
}

func TestMaxDepth(t *testing.T) {
	//case 1
	root := Node{
		Val: 1,
		Children: []*Node{&Node{
			Val: 3,
			Children: []*Node{&Node{
				Val:      5,
				Children: nil,
			}, &Node{
				Val:      6,
				Children: nil,
			}},
		}, &Node{
			Val:      2,
			Children: nil,
		}, &Node{
			Val:      4,
			Children: nil,
		}},
	}

	res := maxDepth(&root)
	assert.Equal(t, 3, res, "Error in case 1")

}
