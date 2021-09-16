package day9_8

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func totalMoney(n int) int {
	res := 0
	week := 0
	monday := 1
	for i:=0; i<n ;i++{
		res += monday+week
		week++
		if week == 7{
			week = 0
			monday++
		}
	}
	return res
}

func TestTotalMoney(t *testing.T){
	//case 1
	res := totalMoney(10)
	assert.Equal(t, 37, res, "Error in case 1")
}