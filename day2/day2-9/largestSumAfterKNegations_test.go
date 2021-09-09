package day2_9

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func largestSumAfterKNegations(nums []int, k int) int {
	for i := 0; i<k; i++{
		target := findLeast(nums)
		nums[target] = 0-nums[target]		//始终将最小的值取反
	}

	res := 0
	for i := range nums{
		res += nums[i]
	}
	return res


}

func findLeast(nums []int)int{	//寻找最小的值
	least := 100
	target := 0
	for i := range nums{
		if nums[i]<least{
			least = nums[i]
			target = i
		}
	}
	return target
}

func Test_largestSumAfterKNegations(t *testing.T){
	//case 1
	nums := []int{4,2,3}
	res := largestSumAfterKNegations(nums,1)
	assert.Equal(t,5, res, "Error in case 1")

	//case 2
	nums = []int{2,-3,-1,5,-4}
	res = largestSumAfterKNegations(nums,2)
	assert.Equal(t, 13, res, "Error in case 2")
}