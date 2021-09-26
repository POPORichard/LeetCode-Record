package day19_9_singleNumber

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func anotherSingleNumber(nums []int)int{
	res := 0
	for i:=0; i<32; i++{
		bit := 0
		for t := range nums{
			bit += (nums[t] >> i) & 1
		}
		res += (bit % 3) << i
	}
	return res
}

func TestAnotherSingleNumber(t *testing.T){
	//case 1
	nums := []int{9,1,7,9,7,9,7}
	res := anotherSingleNumber(nums)
	assert.Equal(t, 1, res, "Error in case 1")
}
