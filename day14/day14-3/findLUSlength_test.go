package day14_3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func findLUSlength(a string, b string) int {
	if a == b {return -1}
	if len(a) > len(b) {return len(a)}
	return len(b)
}

func TestFindLUSlength(t *testing.T){
	//case 1
	res := findLUSlength("aba", "cdc")
	assert.Equal(t, 3, res, "Error in case 1")

}