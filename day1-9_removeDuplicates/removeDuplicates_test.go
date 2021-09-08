package day1_9_removeDuplicates

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func removeDuplicates(s string) string {
	for i:=1;i<len(s);i++{
		if s[i] == s[i-1]{
			s = s[:i-1]+s[i+1:]
			i = 0
		}
	}
	return s
}

// 更好的解决方法
// 用栈的思想
func betterRemoveDuplicates(s string) string{
	t := make([]byte,0,len(s))
	t = append(t,s[0])
	for i:=1;i<len(s);i++{
		if len(t)>0 && s[i] == t[len(t)-1]{
			t= t[:len(t)-1]
		}else{
			t = append(t,s[i])
		}
	}
	return string(t)
}

func Test_removeDuplicates(t *testing.T){
	//case 1
	s := "abbaca"
	res := removeDuplicates(s)
	assert.Equal(t ,"ca" ,res ,"Error in case 1" )
}

func Test_betterRemoveDuplicates(t *testing.T){
	//case 1
	s := "abbaca"
	res := betterRemoveDuplicates(s)
	assert.Equal(t ,"ca" ,res ,"Error in case 1" )
}