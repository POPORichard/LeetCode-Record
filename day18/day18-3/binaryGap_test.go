package day18_3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func binaryGap(n int) int {
	maxLen := 0
	thisLen := 0
	flag := false
	for n > 0 {
		if (n & 1) > 0 {
			if flag && maxLen < thisLen {
				maxLen = thisLen
			}
			thisLen = 0
			flag = true
		}
		if flag {
			thisLen++
		}

		n = n >> 1
	}

	return maxLen
}

func TestBinaryGap(t *testing.T){
	res := binaryGap(22)
	assert.Equal(t, 2,res, "Error in case 1")
}
