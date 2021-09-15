package day8_1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func maxProductDifference(nums []int) int {
	length := len(nums)
	for i:=0; i<length-1; i++{
		biggest := nums[i]
		target := -1
		for t:=i+1; t<length; t++{
			if biggest < nums[t]{
				biggest = nums[t]
				target = t
			}
		}
		if target != -1{
			tmp := nums[i]
			nums[i] = nums[target]
			nums[target] = tmp
		}
	}


	return nums[0]*nums[1]-nums[length-1]*nums[length-2]
}

func TestMaxProductDifference(t *testing.T){

	// case 1
	nums := []int{5,6,2,7,4}
	res := maxProductDifference(nums)
	assert.Equal(t, 34, res, "Error in case 1")
}