package day19_2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func numWaterBottles(numBottles int, numExchange int) int {
	count := numBottles

	for{
		rest := numBottles%numExchange
		numBottles = numBottles/numExchange
		if numBottles == 0{
			return count
		}

		count += numBottles

		numBottles += rest
	}

}

func TestNumWaterBottles(t *testing.T){
	res := numWaterBottles(15,4)
	assert.Equal(t, 19, res, "Error in case 1")
}