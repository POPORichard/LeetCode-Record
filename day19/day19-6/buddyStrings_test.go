package day19_6

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}

	v := [26]int{}
	c := 0
	db := false
	for i := range s {
		v[int(s[i])-'a']++
		// 有重复的
		if v[int(s[i])-'a'] > 1 {
			db = true
		}
	}

	for i := range goal {
		v[int(goal[i])-'a']--

		// 字符数不相同
		if v[int(goal[i])-'a'] < 0 {
			return false
		}

		// 统计不同的位置
		if s[i] != goal[i] {
			c++
		}
		if c > 2 {
			return false
		}
	}

	if c == 0 && db {
		return true
	}
	return c == 2
}

func TestBuddyStrings(t *testing.T){
	//case 1
	s := "aaaaaaabc"
	goal := "aaaaaaacb"
	res := buddyStrings(s, goal)
	assert.Equal(t, true, res, "Error in case 1")

	//case 2
	s = "ab"
	goal = "ab"
	res = buddyStrings(s, goal)
	assert.Equal(t, false, res, "Error in case 2")

	//case 3
	s = "abab"
	goal = "baba"
	res = buddyStrings(s, goal)
	assert.Equal(t, false, res, "Error in case 3")
}