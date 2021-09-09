package day2_3

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)


func reverse(x int) int {
	res := 0
	t := 0
	for x != 0 {
		if res < math.MinInt32/10 || res > math.MaxInt32/10 {
			return 0
		}
		t = x % 10
		x /= 10
		res = res*10 + t
	}
	return res
}

func Test_reverse(t *testing.T){

	//case 1
	x:= 123
	res := reverse(x)
	assert.Equal(t,321,res,"Error in case 1")

	//case 2
	x = -123
	res = reverse(x)
	assert.Equal(t,-321,res,"Error in case 1")


}