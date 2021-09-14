package day7_9_lastRemaining

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func betterLastRemaining(n int)int{
	return  dfs(n,true)
}
func dfs(n int, fromLeft bool) int {
	if n == 1 {
		return 1
	}
	// 如果是奇数的话，从左从右都一样
	if n%2 == 1 {
		return 2 * dfs(n/2, !fromLeft)
	}
	// 如果是偶数,左右消除结果有区分
	if fromLeft {
		return 2 * dfs(n/2, !fromLeft)
	}
	return 2*dfs(n/2, !fromLeft) - 1
}

func TestBetterLastRemaining(t *testing.T){
	//case 1
	res := betterLastRemaining(9)
	assert.Equal(t, 6, res, "Error in case 1")

	//case 2
	res = betterLastRemaining(1)
	assert.Equal(t, 1, res, "Error in case 2")

	//case 3
	res = betterLastRemaining(10000)
	assert.Equal(t, 5974, res, "Error in case 3")
}