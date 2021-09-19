package day12_6

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func minCostClimbingStairs(cost []int) int {
	length := len(cost)
	dp := make([]int, length+1)
	for i := 2; i<length+1; i++{
		dp[i] = min(dp[i - 1] + cost[i - 1], dp[i - 2] + cost[i - 2])
	}
	return dp[length]
}

func min(a,b int) int{
	if a <= b{
		return a
	}
	return b
}

func TestMinCostClimbingStairs(t *testing.T){
	//case 1
	cost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	res := minCostClimbingStairs(cost)
	assert.Equal(t, 6, res , "Error in case 1")

}