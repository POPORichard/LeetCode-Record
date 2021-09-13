package day6_9

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func anotherCountSegments(s string) int {
	var count int = 0
	for i := 0;i < len(s);i++ {
		if (i == 0 || s[i-1] == ' ') && s[i] != ' ' {
			count++
		}
	}
	return count
}

func TestAnotherCountSegments(t *testing.T){
	//case 1
	res := anotherCountSegments("Hello, my name is John")
	assert.Equal(t, 5, res, "Error in case 1")

}
