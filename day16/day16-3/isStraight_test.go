package day16_3

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func isStraight(nums []int) bool {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	zeroNumb := 0
	distance := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == 0 {
			zeroNumb++
		} else if nums[i] == nums[i+1] {
			return false
		} else {
			distance += nums[i+1] - nums[i] - 1
		}
	}

	return distance == 0 || zeroNumb >= distance
}

func TestIsStraight(t *testing.T){
	//case 1
	nums := []int{0,0,1,2,5}
	res := isStraight(nums)
	assert.Equal(t, true, res, "Error in case 1")

	//case 2
	nums = []int{1,2,3,4,5}
	res = isStraight(nums)
	assert.Equal(t, true, res, "Error in case 2")
}