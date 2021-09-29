package day22_8

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func longestWPI(hours []int) int {
	hardDay := make([]int,len(hours))
	for i:= range hours{
		if hours[i] >8{
			hardDay[i] = 1
		}else {
			hardDay[i] = -1
		}
	}

	prefixSum := make([]int, len(hardDay)+1)
	for i, v := range hardDay {
		prefixSum[i+1] = prefixSum[i] + v
	}
	stack := make([]int, 0)
	for i, v := range prefixSum {
		if len(stack) == 0 || prefixSum[stack[len(stack)-1]] > v {
			stack = append(stack, i)
		}
	}
	wpi := 0
	for i := len(prefixSum) - 1; i > wpi; i-- {
		for len(stack) > 0 && prefixSum[i] > prefixSum[stack[len(stack)-1]] {
			w := i - stack[len(stack)-1]
			if wpi < w {
				wpi = w
			}
			stack = stack[:len(stack)-1]
		}
	}
	return wpi
}

func TestLongestWPI(t *testing.T){
	//case 1
	hours := []int{9,9,6,0,6,6,9}
	res := longestWPI(hours)
	assert.Equal(t, 3, res, "Error in case 1")

	//case 2
	hours = []int{6,6,6}
	res = longestWPI(hours)
	assert.Equal(t, 0, res, "Error in case 2")
}
