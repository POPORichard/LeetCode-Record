package day12_2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func heightChecker(heights []int) int {
	count := make([]int, 101)
	for _, h := range heights {
		count[h]++
	}

	ret := 0
	j := 0
	for i := 1; i < len(count); i++ {
		for ; count[i] > 0; count[i]-- {
			if heights[j] != i {
				ret++
			}
			j++
		}
	}
	return ret
}

func TestHeightChecker(t *testing.T){
	heights := []int{1,1,4,2,1,3}
	res := heightChecker(heights)
	assert.Equal(t, 3, res, "Error in case 1")

}