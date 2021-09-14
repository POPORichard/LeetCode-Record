package day7_7

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func kLengthApart(nums []int, k int) bool {
	apart := 0
	start := false
	for i := range nums{
		if nums[i] == 1{
			if apart < k && start{return false}
			start = true
			apart = 0
		}
		if nums[i] == 0{
			apart++
		}
	}
	return true
}

func TestKLengthApart(t *testing.T){
	nums := []int{1,0,0,0,1,0,0,1}
	res := kLengthApart(nums, 2)
	assert.Equal(t, true, res, "Error in case 1")

}
