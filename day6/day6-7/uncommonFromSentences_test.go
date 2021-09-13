package day6_7

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func uncommonFromSentences(s1 string, s2 string) []string {
	//切入字符串组
	arr1 := strings.Fields(s1)
	arr2 := strings.Fields(s2)
	//建立hash
	hashMap := map[string]int{}
	for i := range arr1{
		hashMap[arr1[i]]++
	}
	//对比添加
	for i := range arr2{
		hashMap[arr2[i]]++

	}
	//排除不符合条件
	res := make([]string,0,len(s1))
	for k, v := range hashMap{
		if v == 1{
			res = append(res, k)
		}
	}
	return res
}

func TestUncommonFromSentences(t *testing.T){
	//case 1
	a := "this apple is sweet"
	b := "this apple is sour"
	res := uncommonFromSentences(a,b)
	assert.Equal(t, []string{"sweet", "sour"}, res, "Error in case 1")

	//case 2
	a = "fd kss fd"
	b = "fd fd kss"
	res = uncommonFromSentences(a,b)
	assert.Equal(t, []string{}, res, "Error in case 2")

	//case 3
	a = "fo ly ly"
	b = "fo fo etx"
	res = uncommonFromSentences(a,b)
	assert.Equal(t, []string{"etx"}, res, "Error in case 3")



}
