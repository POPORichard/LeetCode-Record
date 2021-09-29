package day22_1_lengthOfLIS

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


func lengthOfLIS(nums []int) int {
	length := len(nums)
	dp := make([]int,0,length)
	for i:=0; i<length; i++{
		dp = append(dp, 1)
	}

	for i:=1 ;i<length; i++{
		for j:=0; j<i; j++{
			if nums[j] < nums[i]{
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	res := 0
	for i:=0; i<length; i++{
		res = max(res, dp[i])
	}
	return res
}

func max(a,b int)int{
	if a > b{
		return a
	}
	return b
}

func TestLengthOfLIS(t *testing.T){
	//case 1
	nums := []int{10,9,2,5,3,7,101,18}
	res := lengthOfLIS(nums)
	assert.Equal(t, 4, res, "Error in case 1")
}