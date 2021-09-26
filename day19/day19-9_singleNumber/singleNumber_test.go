package day19_9_singleNumber

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func singleNumber(nums []int) int {
	hashMap := map[int]int{}

	for i := range nums{
		hashMap[nums[i]]++
	}

	for k,v := range hashMap{
		if v == 1{
			return k
		}
	}
	return -1
}

func TestSingleNumber(t *testing.T){
	//case 1
	nums := []int{9,1,7,9,7,9,7}
	res := singleNumber(nums)
	assert.Equal(t, 1, res, "Error in case 1")
}
