package day24_4

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func moveZeroes(nums []int) {
	left, right, n := 0, 0, len(nums)
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}

func TestMoveZeroes(t *testing.T){
	//case 1
	nums := []int{0,1,0,3,12}
	moveZeroes(nums)
	assert.Equal(t, []int{1,3,12,0,0}, nums, "Error in case 1")
}