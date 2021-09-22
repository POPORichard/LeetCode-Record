package day15_5

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
func dominantIndex(nums []int) int {
	if len(nums) == 1 {return 0}
	first := 0
	second := -1
	target := 0
	for i := range nums{
		if nums[i] > first{
			second = first
			first = nums[i]
			target = i
			continue
		}
		if nums[i]> second{
			second = nums[i]
		}
	}
	if first >= 2*second{
		return target
	}
	return -1
}

func TestDominantIndex(t *testing.T){
	//case 1
	nums :=[]int{3, 6, 1, 0}
	res := dominantIndex(nums)
	assert.Equal(t, 1, res, "Error in case 1")
}
