package day20_7

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func isPalindrome(s string) bool {
	if s == ""{return true}
	t := make([]int32, 0, 50)
	for _,v := range s{
		if (v >96 && 123 >v) || (v > 47 && 58 > v){
			t = append(t, v)
			continue
		}
		if v >64 && 91 >v{
			t = append(t, v+32)
		}
	}

	if len(t) == 0 || len(t) == 1{
		return true
	}

	left := 0
	right := len(t)-1

	for{
		if t[left] != t[right]{
			return false
		}
		left++
		right--
		if left == right || left >right{
			return true
		}
	}
}

func TestIsPalindrome(t *testing.T){
	//case 1
	res := isPalindrome("A man, a plan, a canal: Panama")
	assert.Equal(t, true, res, "Error in case 1")
}
