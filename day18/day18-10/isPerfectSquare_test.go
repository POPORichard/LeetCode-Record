package day18_10

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func isPerfectSquare(num int) bool {
	if num==1 {
		return true
	}
	l, r := 1, num/2
	for l <= r {
		mid := l + (r-l)>>1
		mid2:=mid*mid
		if mid2 == num {
			return true
		} else if mid2 > num {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return false
}

func TestIsPerfectSquare(t *testing.T){
	//case 1
	res := isPerfectSquare(14)
	assert.Equal(t, false, res, "Error in case 1")
}
