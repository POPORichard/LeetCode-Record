package day5_9

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func frequencySort(nums []int) []int {
	hash := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		hash[nums[i]]++
	}
	sort.Slice(nums, func(i, j int) bool {
		if hash[nums[i]] == hash[nums[j]] {
			return nums[i] > nums[j]
		}
		return hash[nums[i]] < hash[nums[j]]
	})
	return nums
}

func TestFrequencySort(t *testing.T){
	//case 1
	nums := []int{1,1,2,2,2,3}
	res := frequencySort(nums)
	assert.Equal(t, []int{3,1,1,2,2,2}, res, "Error in case 1")
}

