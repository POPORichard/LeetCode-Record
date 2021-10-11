package day26_3

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func smallestRangeI(nums []int, k int) int {
	sort.Ints(nums)
	diff := nums[len(nums)-1] - nums[0]
	res := diff - k*2
	if res >=0 {
		return res
	}
	return 0
}

func TestSmallestRangeI(t *testing.T){
	nums := []int{0,10}
	res := smallestRangeI(nums, 2)
	assert.Equal(t, 6, res, "Error in case 1")
}
