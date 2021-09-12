package day5_7

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func checkOnesSegment(s string) bool {
	if s[0]=='0'{
		return false
	}
	for i:=1;i<len(s);i++ {
		if s[i]=='1'&&s[i-1]=='0' {
			return false
		}
	}
	return true
}

func TestCheckOnesSegment(t *testing.T){
	//case 1
	res := checkOnesSegment("1001")
	assert.Equal(t, false, res, "Error in case 1")
}
