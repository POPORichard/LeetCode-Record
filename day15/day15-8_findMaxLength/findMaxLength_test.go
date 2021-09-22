package day15_8_findMaxLength

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
//time out
//子窜的值等于长度的一半即满足条件

func findMaxLength(nums []int) int {
	sum := 0
	longest := 0
	pre := make([]int,0, len(nums))

	for i := range nums{
		sum += nums[i]
		pre = append(pre, sum)

		for t := (i+1)%2; t<i && longest<i-t ; t +=2{
			tmp := 0
			//若起始的位置上值为1，则需要校准
			if nums[t] == 1{
				tmp = 1
			}
			if sum == pre[t] + (i-t+1)/2 - tmp{
				longest = i-t+1
			}
		}
	}
	return longest
}

func TestFindMaxLength(t *testing.T){
	//case 1
	nums := []int{1,1,1,0,0,0,1,1,0}
	res := findMaxLength(nums)
	assert.Equal(t, 8, res, "Error in case 1")
}