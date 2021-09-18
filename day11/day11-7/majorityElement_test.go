package day11_7

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func majorityElement(nums []int) int {
	target := len(nums)/2
	hashMap := map[int]int{}
	for i := range nums{
		hashMap[nums[i]]++
		if hashMap[nums[i]] > target{
			return nums[i]
		}
	}
	return -1
}

func TestMajorityElement(t *testing.T){
	//case 1
	nums := []int{1, 2, 3, 2, 2, 2, 5, 4, 2}
	res := majorityElement(nums)
	assert.Equal(t, 2, res, "Error in case 1")
}
