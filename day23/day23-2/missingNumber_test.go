package day23_2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func missingNumber(nums []int) int {
	value := 0
	for i := 0; i < len(nums); i++ {
		value ^= nums[i] ^ (i+1)
	}
	return value
}

func TestMissingNumber(t *testing.T){
	//case 1
	nums :=[]int{3,0,1}
	res := missingNumber(nums)
	assert.Equal(t, 2, res, "Error in case 1")
}
