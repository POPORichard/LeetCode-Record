package day4_10

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func isPrefixString(s string, words []string) bool {
	length := len(s)
	word := ""
	for i := range words {
		word += words[i]
		if len(word) == length {
			if word == s {
				return true
			}
			return false
		}
		if len(word) > length {
			return false
		}
	}
	return false
}

func TestIsPrefixString(t *testing.T){
	//case 1
	words := []string{"i","love","leetcode","apples"}
	res := isPrefixString("iloveleetcode", words)
	assert.Equal(t, true, res, "Error in case 1")

}
