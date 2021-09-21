package day14_8

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func largeGroupPositions(s string) [][]int {
	cnt := 1
	ans := make([][]int,0,len(s))

	for i := range s {
		if i == len(s)-1 || s[i] != s[i+1] {
			if cnt >= 3 {
				ans = append(ans, []int{i - cnt + 1, i})
			}
			cnt = 1
		} else {
			cnt++
		}
	}

	return ans
}

func TestLargeGroupPositions(t *testing.T){
	//case 1
	res := largeGroupPositions("abbxxxxzzy")
	assert.Equal(t, [][]int{{3,6}}, res, "Error in case 1")
}

