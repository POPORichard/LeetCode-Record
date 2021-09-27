package day20_4_maxValue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//自底向上
func anotherMaxValue(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	dp := make([][]int,m)
	for i := range dp{
		dp[i] = make([]int,n)
	}

	dp[0][0] = grid[0][0]

	for i:=1; i<m; i++{
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for i:=1; i<n; i++{
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}

	for i := 1; i<m; i++{
		for j := 1; j<n; j++{
			dp[i][j] = max(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
return dp[m-1][n-1]
}


func TestAnotherMaxValue(t *testing.T){
	//case 1
	grid := [][]int{{1,3,1},{1,5,1},{4,2,1}}
	res := anotherMaxValue(grid)
	assert.Equal(t, 12, res, "Error in case 1")
}