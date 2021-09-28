package day21_4

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func findCenter(edges [][]int) int {
	hashMap := map[int]int{}

	for i := range edges {
		hashMap[edges[i][0]]++
		hashMap[edges[i][1]]++
		if i>1{
			break
		}
	}

	for k,v := range hashMap{
		if v > 1{
			return k
		}
	}

	return -1
}

func TestFindCenter(t *testing.T){
	edges := [][]int{{1,2},{5,1},{1,3},{1,4}}
	res := findCenter(edges)
	assert.Equal(t, 1, res, "Error in case 1")
}