package day2_5

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func betterCommonChars(words []string) (ans []string) {
	minFreq := [26]int{}	//建立哈希表,存最小出现次数
	for i := range minFreq {
		minFreq[i] = math.MaxInt64
	}
	for _, word := range words {
		freq := [26]int{}	//建哈希表存当前单词的字母出现次数
		for _, b := range word {
			freq[b-'a']++
		}
		for i, f := range freq {
			if f < minFreq[i] {
				minFreq[i] = f	//比较是否需要更新最小出现次数表
			}
		}
	}
	for i := byte(0); i < 26; i++ {
		for j := 0; j < minFreq[i]; j++ {
			ans = append(ans, string('a'+i))	//查最小次数表并送入答案
		}
	}
	return
}

func Test_betterCommonChars(t *testing.T){
	//case 1
	words:= []string{"bella", "label", "roller"}
	res := betterCommonChars(words)
	assert.Equal(t,[]string{"e","l","l"},res,"Error in case 1")

	//case 2
	words= []string{"cool","lock","cook"}
	res = betterCommonChars(words)
	assert.Equal(t,[]string{"c","o"},res,"Error in case 2")
}
