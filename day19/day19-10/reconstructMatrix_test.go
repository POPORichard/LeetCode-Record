package day19_10

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
func reconstructMatrix(upper int, lower int, colsum []int) [][]int {
	res := make([][]int,2)
	for i := range res{
		tmp := make([]int,len(colsum))
		res[i] = tmp
	}
	times := 0

	for i := range colsum{
		if colsum[i] == 2{
			res[0][i] = 1
			res[1][i] = 1
			times++
		}
	}

	if times > upper || times > lower{
		return nil
	}

	up := times
	low := times
	for i := range colsum{
		if colsum[i] == 1{
			if up < upper{
				res[0][i] = 1
				up++
			}else{
				res[1][i] = 1
				low++
			}
		}
	}
	if low != lower{
		return nil
	}

	return res
}

func TestReconstructMatrix(t *testing.T){
	//case 1
	res := reconstructMatrix(5,5,[]int{2,1,2,0,1,0,1,2,0,1})
	assert.Equal(t, [][]int{{1,1,1,0,1,0,0,1,0,0},{1,0,1,0,0,0,1,1,0,1}}, res, "Error in case 1")

	//case 2
	res = reconstructMatrix(1,4,[]int{2,1,2,0,0,2})
	assert.Equal(t, [][]int(nil), res, "Error in case 1")
}
