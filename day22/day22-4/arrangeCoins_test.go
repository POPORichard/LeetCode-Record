package day22_4

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func arrangeCoins(n int) int {
	if n == 0 {
		return 0
	}
	cur := 0
	res := 0
	for cur <= n {
		res++
		cur += res
	}
	if cur == n {
		return res
	}
	return res-1
}

func TestArrangeCoins(t *testing.T){
	//case 1
	res := arrangeCoins(8)
	assert.Equal(t, 3, res, "Error in case 1")
}
