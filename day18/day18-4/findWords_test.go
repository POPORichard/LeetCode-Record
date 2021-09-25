package day18_4

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func findWords(words []string)(list []string) {
	var s0 = [3]string{"qwertyuiop", "asdfghjkl", "zxcvbnm"}
	for _, word := range words {
		runeS := []rune(strings.ToLower(word)) //排除大小,
	ok:
		for _, v := range s0 {
			if !strings.ContainsRune(v,runeS[0]) {
				continue //排除行
			}
			for _,r := range runeS {
				if !strings.ContainsRune(v,r) {
					break ok //排斥不满条件
				}
			}
			list = append(list, word)
		}
	}
	return list
}

func TestFindWords(t *testing.T){
	res := findWords([]string{"Hello","Alaska","Dad","Peace"})
	assert.Equal(t, []string{"Alaska","Dad"}, res, "Error in case 1")
}
