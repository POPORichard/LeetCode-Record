package day25_2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func pivotIndex(nums []int) int {
	right := 0
	for i := range nums{
		right += nums[i]
	}

	left := 0

	for i := range nums{
		right -= nums[i]
		if right == left{
			return i
		}
		left += nums[i]
	}
	return -1
}

func TestPivotIndex(t *testing.T){
	//case 1
	nums := []int{1,7,3,6,5,6}
	res := pivotIndex(nums)
	assert.Equal(t, 3, res, "Error in case 1")
}
