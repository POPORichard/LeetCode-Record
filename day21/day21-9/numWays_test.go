package day21_9

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func numWays(n int) int {
	a, b := 1, 1
	for i := 1; i <= n; i++ {
		a, b = b, (a+b)%1000000007
	}
	return a
}

func TestNumWays(t *testing.T){
	res := numWays(7)
	assert.Equal(t, 21, res, "Error in case 1")
}
