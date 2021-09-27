package day20_4_maxValue

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

//自上向底
//超时
func maxValue(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	return dp(m-1, n-1, grid)
}

func dp(i,j int,grid [][]int)int{
	if i == 0 && j == 0{
		return grid[0][0]
	}

	if i<0 || j< 0{
		return math.MinInt16
	}


	return max(dp(i-1, j, grid), dp(i, j-1, grid)) + grid[i][j]
}

func max(a,b int)int{
	if a>b{
		return a
	}
	return b
}

func TestMaxValue(t *testing.T){
	//case 1
	grid := [][]int{{1,3,1},{1,5,1},{4,2,1}}
	res := maxValue(grid)
	assert.Equal(t, 12, res, "Error in case 1")
}