package day18_8

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func largestAltitude(gain []int) int {
	high := 0
	maxHigh := 0
	for _, num := range gain {
		high += num
		maxHigh = max(maxHigh, high)
	}
	return maxHigh
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func TestLargestAltitude(t *testing.T){
	//case 1
	gain := []int{-5,1,5,0,-7}
	res := largestAltitude(gain)
	assert.Equal(t, 1, res, "Error in case 1")
}
