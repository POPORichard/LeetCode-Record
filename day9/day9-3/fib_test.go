package day9_3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func fib(n int) int {
	if n == 0{return 0}
	if n == 1{return 1}
	res := [2]int{0,1}
	for i := 1; i<n; i++{
		tmp := res[1]
		res[1] += res[0]
		res[0] = tmp
	}
	return res[1]
}

func TestFib(t *testing.T){
	//case 1
	res := fib(4)
	assert.Equal(t, 3, res, "Error in case 1")
}