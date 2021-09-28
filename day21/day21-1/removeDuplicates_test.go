package day21_1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func removeDuplicates(nums []int) int {
	if len(nums) == 0{return 0}
	count :=0
	pointer :=0

	for i := range nums{
		if nums[i] != nums[pointer]{
			pointer++
			nums[pointer] = nums[i]
			count++

		}
	}
	return count
}

func TestRemoveDuplicates(t *testing.T){
	//case 1
	nums := []int{0,0,1,1,1,2,2,3,3,4}
	res := removeDuplicates(nums)
	assert.Equal(t, []int{0,1,2,3,4}, nums[:res+1], "Error in case 1")
}