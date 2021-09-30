package day23_7

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func coinChange(coins []int, amount int) int {
	hashMap := map[int]int{}

	var dp func(int) int
	dp = func(amount int) int {
		if amount == 0 {
			return 0
		}
		if amount < 0 {
			return -1
		}

		if v, ok := hashMap[amount]; ok {
			return v
		}

		res := 1<<31 - 1
		for i := range coins {
			sub := dp(amount - coins[i])
			if sub == -1 {
				continue
			}
			res = min(res, sub+1)
		}
		if res == 1<<31-1 {
			hashMap[amount] = -1
		} else {
			hashMap[amount] = res
		}
		return hashMap[amount]
	}

	return dp(amount)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func TestCoinChange(t *testing.T){
	//case 1
	coins := []int{1, 2, 5}
	res := coinChange(coins, 11)
	assert.Equal(t, 3, res, "Error in case 1")
}