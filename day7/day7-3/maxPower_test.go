package day7_3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func maxPower(s string) int {
	res := 0
	t:=0
	for i:=1; i<len(s); i++{
		if s[i] == s[i-1]{
			t++
		}else{
			if t>res {res = t}
			t=0
		}
	}
	if t>res {res = t}
	return res+1	//记得+1
}

func TestMaxPower(t *testing.T){
	//case 1
	s := "abbcccddddeeeeedcba"
	res := maxPower(s)
	assert.Equal(t, 5, res, "Error in case 1")

	//case 2
	s = "hooraaaaaaaaaaay"
	res = maxPower(s)
	assert.Equal(t, 11, res, "Error in case 2")
}