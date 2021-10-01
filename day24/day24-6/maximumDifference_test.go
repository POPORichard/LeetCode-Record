package day24_6

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func maximumDifference(nums []int) int {
	ans := 0
	preMin := nums[0]
	for j := 1; j < len(nums); j++ {
		ans = max(ans, nums[j]-preMin)
		preMin = min(preMin, nums[j])
	}
	if ans == 0 {
		ans--
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func TestMaximumDifference(t *testing.T)  {
	//case 1
	nums :=[]int{7,1,5,4}
	res :=maximumDifference(nums)
	assert.Equal(t, 4, res, "Error in case 1")
}
