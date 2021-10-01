package day24_2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func subsetXORSum(nums []int) int {
	s := 0
	r := []int{}
	var dfs func(k int)
	dfs = func(k int) {
		t := 0
		for i := 0; i < len(r); i++ {
			t ^= r[i]
		}
		s += t
		for i := k; i < len(nums); i++ {
			r = append(r, nums[i])
			dfs(i + 1)
			r = r[:len(r)-1]
		}
	}
	dfs(0)
	return s
}

func TestSubsetXORSum(t *testing.T){
	//case 1
	nums := []int{5,1,6}
	res := subsetXORSum(nums)
	assert.Equal(t, 28, res, "Error in case 1")
}