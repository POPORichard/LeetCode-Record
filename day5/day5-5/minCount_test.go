package day5_5

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func minCount(coins []int) int {
	cot := 0
	for _,v := range coins {
		cot += v / 2 + v % 2
	}
	return cot
}

func TestMinCount(t *testing.T){
	//case 1
	coin := []int{2,3,10}
	res := minCount(coin)
	assert.Equal(t, 8, res, "Error in case 1")
}
