package day2_8

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func twoSum(nums []int, target int) []int {
	pointerA := 0
	pointerB := 0
	res := []int{0,0}
	for i:=len(nums)-1;i>=0;i--{
		if nums[i] <= target{
			pointerB = i
			break
		}
	}
	for {
		if nums[pointerA]+nums[pointerB] == target{
			res[0] = nums[pointerA]
			res[1] = nums[pointerB]
			return res
		}
		if nums[pointerA]+nums[pointerB] > target{
			pointerB--
		}
		if nums[pointerA]+nums[pointerB] < target{
			pointerA++
		}
	}
}

func Test_twoSun(t *testing.T){
	//case 1
	nums := []int{2,7,11,15}
	res := twoSum(nums,9)
	assert.Equal(t,[]int{2,7},res,"Error in case 1")

	//case 2
	nums = []int{10,26,30,31,47,60}
	res = twoSum(nums,40)
	assert.Equal(t,[]int{10,30},res,"Error in case 2")

}