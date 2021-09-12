package day5_6

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func decompressRLElist(nums []int) []int {
	res := make([]int, 0, len(nums)*3)
	for i:=1; i<len(nums); i+=2 {
		for t:=0; t<nums[i-1]; t++{
			res = append(res,nums[i])
		}
	}
	return res
}

func TestDecompressRLElist(t *testing.T){
	//case 1
	nums := []int{1,2,3,4}
	res := decompressRLElist(nums)
	assert.Equal(t, []int{2,4,4,4}, res, "Error in case 1")
}
