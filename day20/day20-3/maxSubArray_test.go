package day20_3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func maxSubArray(nums []int) int {
	if len(nums) == 0{
		return nums[0]
	}
	last := nums[0]
	res := last
	for i :=1; i<len(nums); i++{
		last = max(last+nums[i], nums[i])
		if last >res{
			res = last
		}
	}
	return res
}

func max(a, b int)int{
	if a > b{
		return a
	}
	return b
}

func TestMaxSubArray(t *testing.T){
	//case 1
	nums :=[]int{-2,1,-3,4,-1,2,1,-5,4}
	res := maxSubArray(nums)
	assert.Equal(t, 6, res, "Error in case 1")

	//case 2
	nums =[]int{-1,-2}
	res = maxSubArray(nums)
	assert.Equal(t, -1, res, "Error in case 2")
}
