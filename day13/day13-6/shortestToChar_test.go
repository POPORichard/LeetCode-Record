package day13_6

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func shortestToChar(s string, c byte) []int {
	ret := make([]int, len(s))
	for index := 0; index < len(s); index++ {
		// 找到C的位置
		if s[index] == c {
			ret[index] = 0

			// 从C的位置向左遍历,直到到达数组最左或者到达下一个C
			for m := index - 1; m >= 0; m-- {
				if s[m] == c {
					break
				} else {
					dd := ret[m+1] + 1
					if ret[m] == 0 || dd < ret[m] {
						ret[m] = dd
					}
				}
			}

			// 从C的位置向右遍历,直到到达数组最右或者到达下一个C,
			for n := index + 1; n < len(s); n++ {
				if s[n] == c {
					break
				} else {
					dd := ret[n-1] + 1
					if ret[n] == 0 || dd < ret[n] {
						ret[n] = dd
					}
				}
			}
		}
	}
	return ret
}

func TestShortestToChar(t *testing.T){
	//case 1
	res := shortestToChar("loveleetcode", 'e')
	assert.Equal(t, []int{3,2,1,0,1,0,0,1,2,2,1,0}, res, "Error in case 1")

}