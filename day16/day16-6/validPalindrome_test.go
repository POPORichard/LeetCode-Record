package day16_6

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func validPalindrome(s string) bool {
	n := len(s)
	flag := false
	var check func(l, r int) bool
	check = func(l, r int) bool {
		for l < r {
			if s[l] == s[r] {
				l++
				r--
			} else if flag == true {
				return false
			} else {
				flag = true
				return check(l+1, r) || check(l, r-1)
			}
		}
		return true
	}
	return check(0, n-1)

}

func TestValidPalindrome(t *testing.T) {
	//case 1
	res := validPalindrome("abca")
	assert.Equal(t, true, res, "Error in case 1")

	//case 2
	res = validPalindrome("abc")
	assert.Equal(t, false, res, "Error in case 1")

	//case 3
	res = validPalindrome("aguokepatgbnvfqmgmlcupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupuculmgmqfvnbgtapekouga")
	assert.Equal(t, true, res, "Error in case 1")
}
