package day10_7

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func divingBoard(shorter int, longer int, k int) []int {
	if k == 0 {
		return []int{}
	}
	if shorter == longer {
		return []int{shorter * k}
	}
	lengths := make([]int, k + 1)
	for i := 0; i <= k; i++ {
		lengths[i] = shorter * (k - i) + longer * i
	}
	return lengths
}

func TestDivingBoard(t *testing.T){
	//case 1
	res := divingBoard(1,2,3)
	assert.Equal(t, []int{3,4,5,6}, res, "Error in case 1")
}