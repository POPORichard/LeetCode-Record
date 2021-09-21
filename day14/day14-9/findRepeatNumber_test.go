package day14_9

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func findRepeatNumber(nums []int) int {
	hashMap := make(map[int]bool, len(nums))
	b := 0
	for i := 0; i < len(nums); i++ {
		if hashMap[nums[i]] {
			b = nums[i]
			break
		}
		hashMap[nums[i]] = true
	}
	return b
}

func TestFindRepeatNumber(t *testing.T){
	//case 1
	nums := []int{2, 3, 1, 0, 2, 5, 3}
	res := findRepeatNumber(nums)
	assert.Equal(t, 2, res, "Error in case 1")
}