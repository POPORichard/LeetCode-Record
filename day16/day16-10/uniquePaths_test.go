package day16_10

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func uniquePaths(m int, n int) int {
	// f[i][j] 表示i,j到0,0路径数
	f := make([][]int, m)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if f[i] == nil {
				f[i] = make([]int, n)
			}
			f[i][j] = 1
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			f[i][j] = f[i-1][j] + f[i][j-1]
		}
	}
	return f[m-1][n-1]
}

func TestUniquePaths(t *testing.T){
	//case 1
	res := uniquePaths(3,7)
	assert.Equal(t, 28, res, "Error in case 1")
}
