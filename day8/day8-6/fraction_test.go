package day8_6

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
func fraction(cont []int) []int {
	if len(cont) <2 {return []int{cont[0],1}}
	up := 1
	down := cont[len(cont)-1]

	for i:=len(cont)-2; i >= 0; i-- {
		up = cont[i] * down + up
		up, down = down, up
	}
	return []int{down, up}
}

func TestFraction(t *testing.T){
	//case 1
	cont := []int{3, 2, 0, 2}
	res := fraction(cont)
	assert.Equal(t, []int{13,4}, res, "Error in case 1")

	//case 2
	cont = []int{3}
	res = fraction(cont)
	assert.Equal(t, []int{3,1}, res, "Error in case 2")
}