package day15_3

import(
	"github.com/stretchr/testify/assert"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var hashMap = map[int]int{}

func findFrequentTreeSum(root *TreeNode) []int {
	hashMap = map[int]int{}
	dfs(root)

	biggest := 0
	ans := make([]int, 0, 10)
	for val, times := range hashMap{
		if times < biggest{
			continue
		}
		if times > biggest{
			biggest = times
			ans = make([]int, 0, 10)
		}
		ans = append(ans, val)
	}
	return ans
}

func dfs(root *TreeNode) int{
	if root.Right == nil && root.Left == nil{
		hashMap[root.Val]++
		return root.Val
	}

	sum := root.Val

	if root.Left != nil{
		sum += dfs(root.Left)
	}

	if root.Right != nil{
		sum += dfs(root.Right)
	}

	hashMap[sum]++

	return sum

}

func TestFindFrequentTreeSum(t *testing.T){
	//case 1
	root := TreeNode{
		Val:   5,
		Left:  &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   -3,
			Left:  nil,
			Right: nil,
		},
	}
	res := findFrequentTreeSum(&root)
	assert.Equal(t, []int{2,-3,4}, res, "Error in case 1")

	//case 2

	root = TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	res = findFrequentTreeSum(&root)
	assert.Equal(t, []int{1}, res, "Error in case 2")
}
