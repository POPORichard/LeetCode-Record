package day14_4

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func purchasePlans(nums []int, target int) int {
	sort.Ints(nums)
	j := len(nums) - 1
	ans := 0
	for i := 0; i < len(nums); i++ {
		for ; j > i; j-- {
			if nums[i]+nums[j] <= target {
				break
			}
		}
		if j > i {
			ans = ans + (j - i)
		}
	}
	return ans % 1000000007
}

func TestPurchasePlans(t *testing.T){
	//case 1
	nums := []int{2,5,3,6}
	res := purchasePlans(nums, 6)
	assert.Equal(t, 1, res, "Error in case 1")
}
