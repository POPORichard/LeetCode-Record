package day20_2_maxProduct

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
//动态规划
//由于有正负，维护两个数组
//由于状态转移只与前一个数有关，空间优化
func maxProduct(nums []int) int {
	maxF, minF, ans := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		mx, mn := maxF, minF
		maxF = max(mx * nums[i], max(nums[i], mn * nums[i]))
		minF = min(mn * nums[i], min(nums[i], mx * nums[i]))
		ans = max(maxF, ans)
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}


func TestMaxproduct(t *testing.T){
	//case 1
	nums := []int{2,3,-2,4}
	res := maxProduct(nums)
	assert.Equal(t, 6, res, "Error in case 1")

	//case 2
	nums = []int{-2,0,0,-1}
	res = maxProduct(nums)
	assert.Equal(t, 0, res, "Error in case 2")

	//case 3
	nums = []int{1,2,3,4,5,6,7}
	res = maxProduct(nums)
	assert.Equal(t, 5040, res, "Error in case 3")

	//case 4
	nums = []int{0,2}
	res = maxProduct(nums)
	assert.Equal(t, 2, res, "Error in case 4")


	//case 5
	nums = []int{0,0,2}
	res = maxProduct(nums)
	assert.Equal(t, 2, res, "Error in case 5")

	//case 6
	nums = []int{3,-1,4}
	res = maxProduct(nums)
	assert.Equal(t, 4, res, "Error in case 6")

	//case 7
	nums = []int{-2,3,-4}
	res = maxProduct(nums)
	assert.Equal(t, 24, res, "Error in case 7")

}