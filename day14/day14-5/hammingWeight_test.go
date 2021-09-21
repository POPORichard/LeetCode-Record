package day14_5

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func hammingWeight(num uint32) int {
	res := 0
	for ; num > 0; num &= num - 1 {
		res++
	}
	return res
}

func TestHammingWeight(t *testing.T){
	res := hammingWeight(11)
	assert.Equal(t, 3, res, "Error in case 1")
}
