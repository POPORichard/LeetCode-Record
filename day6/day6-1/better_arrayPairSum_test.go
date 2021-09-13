package day6_1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func betterArrayPairSum(nums []int)int{
	res := 0
	for i:=0; i<len(nums)-1; i++{
		target := -1
		largest := nums[i]
		for t:=i+1; t<len(nums); t++{
			if nums[t]> largest {
				largest = nums[t]
				target = t
			}
		}
		if target != -1{
			tem := nums[i]
			nums[i] = nums[target]
			nums[target] = tem
		}

		if i%2 == 1{			//优化：直接在第一次排序时就取出目标值
			res += nums[i]
		}
	}
	res += nums[len(nums)-1]
	return res
}

func TestBetterArrayPairSum(t *testing.T){
	//case 1
	nums := []int{6,2,6,5,1,2}
	res := betterArrayPairSum(nums)
	assert.Equal(t, 9, res, "Error in case 1")
}