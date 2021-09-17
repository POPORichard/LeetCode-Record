package day10_6

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func balancedStringSplit(s string) int {
	count := 0
	result:= 0
	for i:=0;i<len(s);i++{
		if s[i]=='R'{
			count++
		}else if s[i]=='L'{
			count--
		}
		if count==0{
			result++
		}
	}
	return result
}

func TestBalancedStringSplit(t *testing.T){
	//case 1
	res := balancedStringSplit("RLRRLLRLRL")
	assert.Equal(t, 4, res , "Error in case 1")
}
