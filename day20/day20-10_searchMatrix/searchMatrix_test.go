package day20_10_searchMatrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//搜索二维矩阵
//每行的元素从左到右升序排列
//每列的元素从上到下升序排列

func searchMatrix(matrix [][]int, target int) bool {
	row := len(matrix)-1
	col := 0

	for row >= 0 && col < len(matrix[0]){
		if matrix[row][col] > target{
			row--
		}else if matrix[row][col] < target{
			col++
		}else {
			return true
		}
	}
	return false
}

func TestSearchMatrix(t *testing.T){
	//case 1
	matrix := [][]int{{1,4,7,11,15},{2,5,8,12,19},{3,6,9,16,22},{10,13,14,17,24},{18,21,23,26,30}}
	res := searchMatrix(matrix, 5)
	assert.Equal(t, true, res, "Error in case 1")
}

