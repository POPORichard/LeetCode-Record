package day8_8

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func buildArray(nums []int) []int {
	ans := make([]int, len(nums))
	for i := range nums{
		ans[i] = nums[nums[i]]
	}
	return ans
}

func TestBuildArray(t *testing.T){
	//case 1
	nums := []int{0,2,1,5,3,4}
	res := buildArray(nums)
	assert.Equal(t, []int{0, 1, 2, 4, 5, 3}, res, "Error in case 1")

}