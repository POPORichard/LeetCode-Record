package day24_7

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) || s == t {
		return false
	}
	var chars [26]int
	for i := 0; i < len(s); i++ {
		chars[s[i]-'a']++
		chars[t[i]-'a']--
	}
	for _, c := range chars {
		if c != 0 {
			return false
		}
	}
	return true
}

func TestIsAnagram(t *testing.T){
	//case 1
	s := "anagram"
	ta := "nagaram"
	res := isAnagram(s, ta)
	assert.Equal(t, true, res, "Error in case 1")
}