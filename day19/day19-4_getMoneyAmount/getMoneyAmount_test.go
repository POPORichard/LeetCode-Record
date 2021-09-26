package day19_4_getMoneyAmount

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

//unsolved
func getMoneyAmount(n int) int {
	dp := make([][]int,n+1)
	for i:=0 ;i<n+1; i++{
		dp[i] = make([]int, n+1)
	}

	for len :=2; len<=n; len++{
		for start :=1; start<= n-len+1; start++{
			minres := math.MaxInt32
			for piv := start + (len - 1)/2; piv< start +len-1; piv++{
				res := piv + max(dp[start][piv-1], dp[piv+1][start + len - 1])
				minres = min(res, minres)
			}
			dp[start][start + len -1] = minres
		}
	}
	return dp[1][n]
}

func max (a,b int) int{
	if a > b{
		return a
	}
	return b
}

func min(a,b int) int{
	if a > b{
		return b
	}
	return a
}

func TestGetMoneyAmount(t *testing.T){
	res := getMoneyAmount(10)
	assert.Equal(t, 16, res, "Error in case 1")
}