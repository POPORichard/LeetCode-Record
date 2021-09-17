package day10_10

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
func hasAlternatingBits(n int) bool {
	for n != 0 {
		if (n & 1) == ((n >> 1) & 1) {
			return false
		}
		n >>= 1
	}
	return true
}

func TestHasAlternatingBits(t *testing.T){
	res := hasAlternatingBits(10)
	assert.Equal(t, true, res, "Error in case 1")
}