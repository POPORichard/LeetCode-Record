package day26_6

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func maxProfit(prices []int, fee int) int {
	length := len(prices)

	dp := make([][]int, length)
	dp[0] = []int{0, -prices[0]-fee}
	for i := 1; i<length; i++{
		dp[i] = []int{math.MinInt32,math.MaxInt32}
	}
	for i := 1; i<length; i++{
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i] - fee)
	}
	return dp[length-1][0]
}

func max(a,b int)int{
	if a > b {
		return a
	}
	return b
}

func TestMaxProfit(t *testing.T){
	//case 1
	prices := []int{1, 3, 2, 8, 4, 9}
	res := maxProfit(prices, 2)
	assert.Equal(t, 8, res, "Error in case 1")
}