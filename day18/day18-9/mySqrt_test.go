package day18_9

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func mySqrt(x int) int {
	l, r := 0, x
	ans := -1
	for l <= r {
		mid := l + (r - l) / 2
		if mid * mid <= x {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ans
}

func TestMySqrt(t *testing.T){
	//case 1
	res := mySqrt(8)
	assert.Equal(t, 2, res, "Error in case 1")
}
