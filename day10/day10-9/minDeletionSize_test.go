package day10_9

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
func minDeletionSize(strs []string) int {
	res := 0
	for i := range strs[0] {
		for t := 1; t < len(strs); t++ {
			if strs[t-1][i] > strs[t][i] {
				res++
				break
			}
		}
	}
	return res
}

func TestMinDeletionSize(t *testing.T){
	//case 1
	strs := []string{"cba","daf","ghi"}
	res := minDeletionSize(strs)
	assert.Equal(t, 1, res, "Error in case 1")

	//case 2
	strs = []string{"a","b"}
	res = minDeletionSize(strs)
	assert.Equal(t, 0, res, "Error in case 1")

}
