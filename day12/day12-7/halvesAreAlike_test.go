package day12_7

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func halvesAreAlike(s string) bool {
	set := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
		'A': true,
		'E': true,
		'I': true,
		'O': true,
		'U': true,
	}
	delta := 0
	for i := 0; i < len(s)/2; i++ {
		if _, ok := set[s[i]]; ok {
			delta++
		}
		if _, ok := set[s[len(s)-1-i]]; ok {
			delta--
		}
	}
	return delta == 0
}

func TestHalvesAreAlike(t *testing.T){
	res := halvesAreAlike("AbCdEfGh")
	assert.Equal(t, true, res, "Error in case 1")
}
