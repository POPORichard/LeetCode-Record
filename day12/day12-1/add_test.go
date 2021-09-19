package day12_1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func add(a int, b int) int {
	for b != 0 {
		XOR := a ^ b
		And := (a & b) << 1
		a = XOR
		b = And
	}
	return a
}

func TestAdd(t *testing.T){
	//case 1
	res := add(1,3)
	assert.Equal(t, 4, res, "Error in case 1")
}
