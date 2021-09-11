package day4_6

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func getRow(rowIndex int) []int {
	res := make([][]int, rowIndex+1)
	for i := range res {
		res[i] = make([]int, i+1)
		res[i][0], res[i][i] = 1, 1
		for j := 1; j < i; j++ {
			res[i][j] = res[i-1][j-1] + res[i-1][j]
		}
	}
	return res[rowIndex]
}

func TestGetRow(t *testing.T){
	//case 1
	res := getRow(3)
	assert.Equal(t, []int{1,3,3,1}, res, "Error in case 1")
}