package day9_5

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func numOfStrings(patterns []string, word string) int {
	res := 0
	for _,w := range patterns{
		if strings.Contains(word, w){
			res++
		}
	}
	return res
}

func TestNumOfStrings(t *testing.T){
	//case 1
	patterns := []string{"a","abc","bc","d"}
	res := numOfStrings(patterns, "abc")
	assert.Equal(t, 3, res, "Error in case 1")
}