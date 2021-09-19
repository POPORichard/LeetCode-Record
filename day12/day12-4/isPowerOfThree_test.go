package day12_4

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"strings"
	"testing"
)

func isPowerOfThree(n int) bool {
	if n < 1 {
		return false
	}
	s := strconv.FormatInt(int64(n), 3)
	return s[0:1] == "1" && strings.Count(s, "0") == len(s)-1
}

func TestIsPowerOfThree(t *testing.T){
	//case 1
	res := isPowerOfThree(27)
	assert.Equal(t, true, res, "Error in case 1")
}
