package day15_2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
func printNumbers(n int) []int {
	target := 1
	for i := 0; i < n; i++ {
		target = target * 10
	}
	ans := make([]int, 0, target)
	for i := 1; i < target; i++ {
		ans = append(ans, i)
	}
	return ans
}

func TestPrintNumbers(t *testing.T){
	//case 1
	res := printNumbers(2)
	assert.Equal(t, []int{1,2,3,4,5,6,7,8,9}, res, "Error in case 1")
}
