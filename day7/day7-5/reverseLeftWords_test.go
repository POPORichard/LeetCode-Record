package day7_5

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
func reverseLeftWords(s string, n int) string {
	tmp := s[:n]
	s = s[n:]+tmp
	return s
}

func TestReverseLeftWords(t *testing.T){
	//case 1
	s := "lrloseumgh"
	res := reverseLeftWords(s, 6)
	assert.Equal(t, "umghlrlose", res, "Error in case 1")
}