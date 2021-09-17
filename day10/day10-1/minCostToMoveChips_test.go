package day10_1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func minCostToMoveChips(position []int) int {
	odd := 0
	for i := range position{
		if position[i]%2 == 1{
			odd++
		}
	}
	if odd > len(position)-odd {
		return len(position)-odd
	}
	return odd
}

func TestMinCostToMoveChips(t *testing.T){
	//case 1
	chips := []int{2,2,2,3,3}
	res := minCostToMoveChips(chips)
	assert.Equal(t, 2, res, "Error in case 1")
}