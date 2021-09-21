package day14_10

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

func TestLongestCommonPrefix(t *testing.T){
	//case 1
	strs := []string{"flower","flow","flight"}
	res := longestCommonPrefix(strs)
	assert.Equal(t, "fl", res ,"Error in case 1")
}
