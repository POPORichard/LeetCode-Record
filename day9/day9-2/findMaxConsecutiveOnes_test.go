package day9_2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
func findMaxConsecutiveOnes(nums []int) int {
	res := 0
	cur := 0
	for i := range nums{
		if nums[i] == 1{
			cur++
		}else {
			if cur > res{res = cur}
			cur = 0
		}
	}
	if cur > res{return cur}
	return  res
}

func TestFindMaxConsecutiveOnes(t *testing.T){
	//case 1
	num := []int{1,1,0,1,1,1}
	res := findMaxConsecutiveOnes(num)
	assert.Equal(t, 3, res, "Error in case 1")

}