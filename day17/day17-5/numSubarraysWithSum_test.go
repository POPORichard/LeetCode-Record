package day17_5

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func numSubarraysWithSum(nums []int, goal int) int {
	hashMap := map[int]int{0:1}
	count := 0
	sum := 0
	for i := range nums{
		sum += nums[i]
		count += hashMap[sum - goal]
		hashMap[sum]++
	}
	return count
}

func TestNumSubarraysWithSum(t *testing.T){
	nums := []int{1,0,1,0,1}
	res := numSubarraysWithSum(nums, 2)
	assert.Equal(t, 4, res, "Error in case 1")
}