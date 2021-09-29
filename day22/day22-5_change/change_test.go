package day22_5_change

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}
	return dp[amount]
}

func TestChange(t *testing.T){
	coins := []int{1,2,5}
	res := change(5,coins)
	assert.Equal(t, 4, res, "Error in case 1")
}
