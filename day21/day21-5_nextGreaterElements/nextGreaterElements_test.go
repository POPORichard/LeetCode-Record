package day21_5_nextGreaterElements

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func nextGreaterElements(nums []int) []int {
	length := len(nums)
	res := make([]int, len(nums))
	stack := make([]int,0,len(nums))
	stack = append(stack, nums[len(nums)-1])

	for i := (length-1)*2; i>=0; i--{
		for len(stack)!=0 && stack[len(stack)-1] <= nums[i%length]{
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0{
			res[i%length] = -1
		}else {
			res[i%length] = stack[len(stack)-1]
		}
		stack = append(stack, nums[i%length])
	}

	return res
}

func TestNextGreaterElements(t *testing.T){
	//case 1
	nums := []int{1,2,1}
	res := nextGreaterElements(nums)
	assert.Equal(t, []int{2,-1,2}, res, "Error in case 1")
}