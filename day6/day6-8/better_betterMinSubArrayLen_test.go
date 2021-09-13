package day6_8

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func betterMinSubArrayLen(target int, nums []int) int {
	n := len(nums)
	res := 0
	sum := 0
	l, r := 0, 0
	for l <= r && r < n {
		sum += nums[r]
		for l <= r && sum >= target {
			sum -= nums[l]
			if  res == 0 {
				res = r-l+1
			}else if res > r-l+1 {
				res = r-l+1
			}
			l++
		}
		r++
	}
	return res

}


func TestBetterMinSubArrayLen(t *testing.T){
	//case 1
	nums := []int{2,3,1,2,4,3}
	res := betterMinSubArrayLen(7, nums)
	assert.Equal(t, 2, res, "Error in case 1")

	//case 2
	nums = []int{1,4,4}
	res = betterMinSubArrayLen(4, nums)
	assert.Equal(t, 1, res, "Error in case 2")

	//case 3
	nums = []int{1,1,1,1,1,1,1,1}
	res = betterMinSubArrayLen(11, nums)
	assert.Equal(t, 0, res, "Error in case 3")

	//case 4
	nums = []int{10,2,3}
	res = betterMinSubArrayLen(6, nums)
	assert.Equal(t, 1, res, "Error in case 4")

	//case 5
	nums = []int{1,2,3,4,5}
	res = betterMinSubArrayLen(11, nums)
	assert.Equal(t, 3, res, "Error in case 5")

	//case 6
	nums = []int{5,1,3,5,10,7,4,9,2,8}
	res = betterMinSubArrayLen(15, nums)
	assert.Equal(t, 2, res, "Error in case 6")


	//case 7
	nums = []int{2,16,14,15}
	res = betterMinSubArrayLen(20, nums)
	assert.Equal(t, 2, res, "Error in case 7")


}