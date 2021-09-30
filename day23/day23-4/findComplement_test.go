package day23_4

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func findComplement(num int) int {
	res := 0
	t := 0

	for num != 0{
		res += (num&1 ^ 1)<<t
		num >>= 1
		t++
	}
	return res
}

func TestFindComplement(t *testing.T){
	//case 1
	res := findComplement(5)
	assert.Equal(t, 2, res, "Error in case 1")
}