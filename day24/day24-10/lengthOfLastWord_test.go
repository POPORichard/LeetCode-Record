package day24_10

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func lengthOfLastWord(s string) int {
	s= strings.TrimSpace(s)
	arr := strings.Split(s, " ")
	n := len(arr)
	return len(arr[n-1])
}

func TestLengthOfLastWord(t *testing.T){
	//case 1
	res := lengthOfLastWord("   fly me   to   the moon  ")
	assert.Equal(t, 4, res, "Error in case 1")

}
