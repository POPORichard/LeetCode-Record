package day16_9_subarraysDivByK

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
//超时
func subarraysDivByK(nums []int, k int) int {
	sum := 0
	count := 0
	hashMap := map[int]int{0: 1}

	for i := range nums{
		sum += nums[i]


		for n,v := range hashMap{
			if (sum - n)%k == 0{
				count += v
			}
		}
		hashMap[sum]++
	}

	return count
}

func TestSubarraysDivByK(t *testing.T){
	//case 1
	nums := []int{4,5,0,-2,-3,1}
	res := subarraysDivByK(nums, 5)
	assert.Equal(t, 7, res, "Error in case 1")
}
