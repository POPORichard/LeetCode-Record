package day22_2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func transpose(matrix [][]int) [][]int {
	m := len(matrix)
	n := len(matrix[0])

	res := make([][]int,n)
	for i:=0; i<n; i++{
		tmp := make([]int,m)
		res[i] = append(res[i], tmp...)
	}

	for i:=0; i<m; i++{
		for j:=0; j<n; j++{
			res[j][i] = matrix[i][j]
		}
	}
	return res
}

func TestTranspose(t *testing.T){
	//case 1
	matrix := [][]int{{1,2,3},{4,5,6},{7,8,9}}
	res := transpose(matrix)
	assert.Equal(t, [][]int{{1,4,7},{2,5,8},{3,6,9}}, res, "Error in case 1")

	//case 2
	matrix = [][]int{{1,2,3},{4,5,6}}
	res = transpose(matrix)
	assert.Equal(t, [][]int{{1,4},{2,5},{3,6}}, res, "Error in case 1")
}
