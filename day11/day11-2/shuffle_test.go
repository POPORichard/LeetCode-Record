package day11_2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func shuffle(nums []int, n int) []int {
	ans := make([]int,0,2*n)
	for i:= 0; i<n ;i++{
		ans = append(ans,nums[i])
		ans = append(ans,nums[i+n])
	}
	return ans
}

func TestShuffle(t *testing.T){
	//case 1
	nums := []int{1,2,3,4,4,3,2,1}
	res := shuffle(nums,4)
	assert.Equal(t, []int{1,4,2,3,3,2,4,1}, res, "Error in case 1")
}