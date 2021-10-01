package day24_9

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func maxLengthBetweenEqualCharacters(s string) int {
	ans := -1
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] && j-i-1 > ans {
				ans = j - i - 1
			}
		}
	}
	return ans
}

func TestMaxLengthBetweenEqualCharacters(t *testing.T){
	res := maxLengthBetweenEqualCharacters("cabbac")
	assert.Equal(t, 4, res, "Error in case 1")
}
