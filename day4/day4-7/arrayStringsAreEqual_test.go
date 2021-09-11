package day4_7

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	w1 := ""
	w2 := ""
	for i:=0; i<len(word1); i++{
		w1 += word1[i]
	}
	for i:=0; i<len(word2); i++{
		w2 += word2[i]
	}

	if w1 == w2 {return true}
	return false
}

func TestArrayStringsAreEqual(t *testing.T){

	//case 1
	word1 :=[]string{"a", "cb"}
	word2 :=[]string{"ab", "c"}
	res := arrayStringsAreEqual(word1, word2)
	assert.Equal(t, false, res, "Error in case 1")
}
