package day6_10

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func checkPerfectNumber(num int) bool {
	if num == 1 {return false}
	res := 1
	for i:=2;i*i<num;i++{
		if num%i == 0 {
			res += i
			if i*i != num{	//避免重复
				res += num/i
			}
		}
	}
	return res == num
}

func TestCheckPerfectNumber(t *testing.T){
	res := checkPerfectNumber(28)
	assert.Equal(t, true, res, "Error in case 1")
}