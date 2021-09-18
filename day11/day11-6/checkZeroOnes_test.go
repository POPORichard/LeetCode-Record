package day11_6

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
func checkZeroOnes(s string) bool {
	zero := 0
	one := 0
	z := 0
	o := 0
	for i := range s{
		if s[i] == 48{
			z++
			if o > one{
				one = o
			}
			o = 0
		}else {
			o++
			if z > zero{
				zero = z
			}
			z = 0
		}
	}
	if o > one{
		one = o
	}
	if z > zero{
		zero = z
	}
	if one > zero{
		return true
	}
	return false
}

func TestCheckZeroOnes(t *testing.T){
	//case 1
	res := checkZeroOnes("110100010")
	assert.Equal(t, false, res, "Error in case 1")
}
