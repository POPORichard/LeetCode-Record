package day2_7

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] + nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}


func Test_maxSubArray(t *testing.T){
	//case 1
	nums := []int{-2,1,-3,4,-1,2,1,-5,4}
	res := maxSubArray(nums)
	assert.Equal(t,6,res,"Error in case 1")

	//case 2
	nums = []int{-1,-2}
	res = maxSubArray(nums)
	assert.Equal(t,-1,res,"Error in case 2")
}