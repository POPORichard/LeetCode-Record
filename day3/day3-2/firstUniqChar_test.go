package day3_2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func firstUniqChar(s string) byte {
	if len(s) == 0 {return ' '}
	hash := make([]int,26,26)
	hashMap := make(map[uint8]int)
	for i:= range s{
		hash[s[i]-'a']++
		hashMap[s[i]]=i
	}
	res := make([]uint8,0,26)
	for  i := range hash{
		if hash[i] == 1{
			res = append(res,uint8(i+'a'))
		}
	}
	if len(res) == 1 {return res[0]}
	if len(res) == 0 {return ' '}
	ans := res[0]
	for i :=range res{
		if hashMap[res[i]] < hashMap[ans]{
			ans = res[i]
		}
	}
	return ans
}

func TestFirstUniqChar(t *testing.T){
	//case 1
	s := "abaccdeff"
	res := firstUniqChar(s)
	assert.Equal(t, uint8('b'), res, "Error in case 1")

	//case 2

}