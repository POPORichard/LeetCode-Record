package day4_9

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func countConsistentStrings(allowed string, words []string) int {
	hash :=make([]bool,26,26)	//建立hash
	for i := range allowed{
		hash[allowed[i]-'a'] = true
	}

	res := 0

	Loop:
	for i := range words{
		flag := false
		for t:= range words[i]{
			 flag = hash[words[i][t]-'a']		//判断是否满足条件
			 if flag != true{
			 	continue Loop
			 }
		}
		res ++
	}
	return res
}

func TestCountConsistentStrings(t *testing.T){
	//case 1
	words := []string{"ad","bd","aaab","baa","badab"}
	res := countConsistentStrings("ab", words)
	assert.Equal(t, 2, res, "Error in case 1" )


	//case 2
	words = []string{"a","b","c","ab","ac","bc","abc"}
	res = countConsistentStrings("abc", words)
	assert.Equal(t, 7, res, "Error in case 2" )
}
