package day16_1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func minArray(numbers []int) int {
	length := len(numbers)
	if length == 1 {return numbers[0]}
	min := numbers[length-1]
	for i:=length-2; i>=0; i--{
		if numbers[i] <= min {
			min = numbers[i]
		}else {
			return min
		}
	}
	return min
}

func TestMinArray(t *testing.T){
	//case 1
	numbers := []int{3,4,5,1,2}
	res := minArray(numbers)
	assert.Equal(t, 1, res, "Error in case 1")
}