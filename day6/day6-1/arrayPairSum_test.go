package day6_1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func arrayPairSum(nums []int) int {
	//先排序
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
	}

	//取出奇数位上的数加起来即为答案
	res := 0
	for i:=1; i<len(nums); i+=2{
		res += nums[i]
	}
	return res
}

//该方法超时
func TestArrayPairSum(t *testing.T){
	//case 1
	nums := []int{6,2,6,5,1,2}
	res := arrayPairSum(nums)
	assert.Equal(t, 9, res, "Error in case 1")
}



