package day5_3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func isPrefixOfWord(sentence string, searchWord string) int {
	wi := 0
	si := 0

	woodcut := 0
	for si < len(sentence) {
		woodcut++
		for si < len(sentence) {
			if sentence[si] == searchWord[wi] {
				si++
				wi++
				if wi == len(searchWord) {
					return woodcut
				}
			} else {
				break
			}
		}
		for ; si < len(sentence); si++ {
			if sentence[si] == ' ' {
				si++
				wi = 0
				break
			}
		}
	}
	return -1
}

func TestIsPrefixOfWord(t *testing.T){
	//case 1
	res := isPrefixOfWord("i love eating burger", "burg")
	assert.Equal(t, 4, res, "Error in case 1")
}
