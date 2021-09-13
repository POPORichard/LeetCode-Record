package day6_3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func checkIfPangram(sentence string) bool {
	//建立hash映射
	hashMap := make([]int,26,26)
	for i := range sentence{
		hashMap[sentence[i]-'a']++
	}
	//看hash表上是否每位都至少为1
	for i := range hashMap {
		if hashMap[i]<1{
			return false
		}
	}
	return true
}

func TestCheckIfPangram(t *testing.T){
	//case 1
	res := checkIfPangram("thequickbrownfoxjumpsoverthelazydog")
	assert.Equal(t, true, res, "Error in case 1")

	//case 2
	res = checkIfPangram("leetcode")
	assert.Equal(t, false, res, "Error in case 1")

}