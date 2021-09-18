package _1_4

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func largestPerimeter(nums []int) int {
	sort.Ints(nums)
	for i := len(nums) - 1; i >= 2; i-- {
		if nums[i-2]+nums[i-1] > nums[i] {
			return nums[i-2] + nums[i-1] + nums[i]
		}
	}
	return 0
}

func TestLargestPerimeter(t *testing.T){
	//case 1
	nums := []int{3,6,2,3}
	res := largestPerimeter(nums)
	assert.Equal(t, 8, res, "Error in case 1")
}
