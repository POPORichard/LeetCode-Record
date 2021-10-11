package day26_2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func detectCapitalUse(word string) bool {
	upperFound := false
	lowerFound := false
	for i := len(word) - 1; i >= 0; i-- {
		if word[i] >= 'A' && word[i] <= 'Z' {
			if upperFound && lowerFound {
				return false
			}
			upperFound = true
		} else {
			if upperFound {
				return false
			}
			lowerFound = true
		}
	}
	return true
}

func TestDetectCapitalUse(t *testing.T){
	//case 1
	res := detectCapitalUse("Google")
	assert.Equal(t, true, res, "Error in case 1")
}