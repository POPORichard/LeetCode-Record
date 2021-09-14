package day7_9_lastRemaining

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//本方法超时

func lastRemaining(n int) int {
	if n ==1 {return 1}
	if n%2 == 1 {n = n-1}
	nums := make([]int,0,n/2+1)
	for i := 2; i<=n; i+=2{
		nums = append(nums, i)
	}
	for {
		if len(nums) == 1 {
			return nums[0]
		}
		for i := len(nums) - 1; i >= 0; i -= 2 {
			nums = append(nums[:i], nums[i+1:]...)
		}
		if len(nums) == 1 {
			return nums[0]
		}
		for i := 0; i < len(nums)-1/2; i++ {
			nums = append(nums[:i], nums[i+1:]...)
		}

	}
}



func TestLastRemaining(t *testing.T){
	//case 1
	res := lastRemaining(9)
	assert.Equal(t, 6, res, "Error in case 1")

	//case 2
	res = lastRemaining(1)
	assert.Equal(t, 1, res, "Error in case 2")

	//case 3
	res = lastRemaining(10000)
	assert.Equal(t, 1, res, "Error in case 3")
}