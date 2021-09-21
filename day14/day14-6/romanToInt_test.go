package day14_6

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func romanToInt(s string) int {

	symbolValues := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000}

	ans := 0

	n := len(s)
	for i := range s {
		value := symbolValues[s[i]]
		if i < n-1 && value < symbolValues[s[i+1]] {
			ans -= value
		} else {
			ans += value
		}
	}
	return ans
}

func TestRomanToInt(t *testing.T){
	//case 1
	res := romanToInt("III")
	assert.Equal(t, 3, res, "Error in case 1")
}
