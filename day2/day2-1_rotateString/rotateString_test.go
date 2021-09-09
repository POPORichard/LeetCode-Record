package day2_1_rotateString

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func rotateString(s string, goal string) bool {
	if len(s)== len(goal) && strings.Contains(s+s, goal){return true}
	return false
}

func Test_rotateString(t *testing.T){
	//case 1
	A := "abcde"
	B := "cdeab"
	res := rotateString(A,B)
	assert.Equal(t,true,res,"Error in case 1")

	//case 2
	A = "abcde"
	B = "abced"
	res = rotateString(A,B)
	assert.Equal(t,false,res,"Error in case 2")

	//case 3
	A = "ohbrwzxvxe"
	B = "uornhegseo"

	res = rotateString(A,B)
	assert.Equal(t,false,res,"Error in case 3")

	//case 4
	A = "bbbacddceeb"
	B = "ceebbbbacdd"
	res = rotateString(A,B)
	assert.Equal(t,true,res,"Error in case 4")
}