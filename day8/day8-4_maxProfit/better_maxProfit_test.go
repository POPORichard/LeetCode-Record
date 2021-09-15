package day8_4_maxProfit

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
//动态规划
func betterMaxProfit(prices []int)int{
	//建立状态
	dp := make([][]int,len(prices))
	for i := range dp{
		dp[i] = make([]int,2)
	}

	//设定初始状态
	//dp[i][0]即第i天且没有持有股票时手里的钱
	//dp[i][1]即第i天且持有股票时手里的钱
	dp[0][0] = 0
	dp[0][1] = -prices[0]

	for i:=1; i<len(prices); i++{
		dp[i][0] = max(dp[i-1][0],dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1],-prices[i])
	}
	//最后一天必定已经卖出
	return dp[len(prices)-1][0]

}

func max(a,b int)int{
	if a>b{
		return a
	}
	return b
}

func TestBetterMaxProfit(t *testing.T){
	//case 1
	prices := []int{7,1,5,3,6,4}
	res := betterMaxProfit(prices)
	assert.Equal(t, 5, res, "Error in case 1")
}
