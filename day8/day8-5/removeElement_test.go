package day8_5

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
func removeElement(nums []int, val int) int {
	left, right := 0, len(nums)
	for left < right {
		if nums[left] == val {
			nums[left] = nums[right-1]
			right--
		} else {
			left++
		}
	}
	return left
}
func TestRemoveElement(t *testing.T){
	//case 1
	nums := []int{0,1,2,2,3,0,4,2}
	res := removeElement(nums,2)
	assert.Equal(t,5, res, "Error in case 1")
}
