package day5_2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func lastRemaining(n int, m int) int {
	 f := 0
	for  i := 2; i != n + 1;i++ {
		f = (m + f) % i
	}
	return f
}

func TestLastRemaining(t *testing.T){
	//case 1
	res := lastRemaining(5,3)
	assert.Equal(t, 3, res , "Error in case 1")
}

