package day1_6

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func thousandSeparator(n int) string {
	s := strconv.Itoa(n)
	length := len(s)+len(s)/3
	add := len(s)/3
	res := make([]uint8,length,length)
	jumpOver := 0
	t := 0
	for i:= len(s)-1;i>=0;i--{
		if t == 3{
			if i-jumpOver == length{
				break
			}
			res[i+add-jumpOver] = 46
			jumpOver++
			t = 0
		}
		res[i+add-jumpOver] = s[i]
		t++
	}
	if res[0] == 0{return string(res[1:])}
	return string(res)
}

func Test_thousandSeparator(t *testing.T){
	//case 1
	n := 987
	res := thousandSeparator(n)
	assert.Equal(t,"987",res,"Error in case 1")

	//case 2
	n = 123456789
	res = thousandSeparator(n)
	assert.Equal(t,"123.456.789",res,"Error in case 2")

	//case 3
	n = 1234
	res = thousandSeparator(n)
	assert.Equal(t,"1.234",res,"Error in case 2")



}