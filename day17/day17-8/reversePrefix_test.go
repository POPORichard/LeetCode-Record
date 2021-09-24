package day17_8

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
func reversePrefix(word string, ch byte) string {
	stack := make([]byte,0,len(word))
	i:=0
	for i = range word{
		if word[i] == ch{
			stack = append(stack, word[i])
			break
		}
		stack = append(stack,word[i])
	}
	if i == len(word)-1 && word[len(word)-1] != ch{
		return word
	}

	tmp := make([]byte,0,len(stack))
	for i:= len(stack)-1; i>=0; i--{
		tmp = append(tmp,stack[i])
	}

	return string(tmp)+word[i+1:]
}

func TestReversePrefix(t *testing.T){
	//case 1
	res := reversePrefix("abcdefd", 'd')
	assert.Equal(t, "dcbaefd", res, "Error in case 1")

}

