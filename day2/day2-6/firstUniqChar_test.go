package day2_6

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func firstUniqChar(s string) int {
	var hash [26]int
	for _,char := range s{	//第一次遍历：建立映射
		hash[char-'a']++
	}
	for i := range s{	//第二次遍历：查找第一个满足条件
		if hash[s[i]-'a'] == 1{
			return i
		}
	}
	return -1
}

func Test_firstUniqChar(t *testing.T){

	//case 1
	res := firstUniqChar("leetcode")
	assert.Equal(t,0,res,"Error in case 1")
}