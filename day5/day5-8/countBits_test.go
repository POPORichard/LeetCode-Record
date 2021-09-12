package day5_8

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func countBits(n int) []int {
	ret := make([]int, n+1)
	for i := 1; i <= n; i++ {
		ret[i] = ret[i>>1]	//相当于除以2
		if i%2 == 1 {
			ret[i]++
		}
	}
	return ret
}

func TestCountBits(t *testing.T){
	//case 1
	ret := countBits(5)
	assert.Equal(t, []int{0,1,1,2,1,2}, ret ,"Error in case 1")
}


