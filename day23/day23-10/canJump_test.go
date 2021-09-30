package day23_10

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func canJump(nums []int) bool {
	farthest := 0
	for i:= 0; i<len(nums)-1; i++{
		farthest = max(farthest, i+nums[i])
		if farthest <= i{
			return false
		}
	}
	return farthest >= len(nums)-1
}

func max(a,b int) int{
	if a>b{
		return a
	}
	return b
}

func TestCanJump(t *testing.T){
	nums := []int{2,3,1,1,4}
	res := canJump(nums)
	assert.Equal(t, true, res, "Error in case 1")
}
