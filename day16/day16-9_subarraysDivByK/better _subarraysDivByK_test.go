package day16_9_subarraysDivByK

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//同余定理
func betterSubarraysDivByK(nums []int, k int) int {
	record := map[int]int{0: 1}
	sum, ans := 0, 0
	for _, elem := range nums {
		sum += elem
		//不能直接用sum%k, 可能有负数的情况
		modulus := (sum % k + k) % k
		ans += record[modulus]
		record[modulus]++
	}
	return ans
}

func TestBetterSubarraysDivByK(t *testing.T){
	//case 1
	nums := []int{4,5,0,-2,-3,1}
	res := betterSubarraysDivByK(nums, 5)
	assert.Equal(t, 7, res, "Error in case 1")
}
