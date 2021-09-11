package day4_3

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func isPalindrome(x int) bool {
	if x<0 {return false}
	if x<10 {return true}
	arr := make([]int,0,100)
	for{
		if x<10 {break}
		tmp := x%10
		arr = append(arr, tmp)
		x = x/10
	}
	arr = append(arr, x)
	for i := 0; i<len(arr)/2; i++{
		if arr[i] != arr[len(arr)-1-i]{
			return false
		}
	}
	return true
}

func TestIsPalindrome(t *testing.T){
	//case 1
	assert.Equal(t, true, isPalindrome(121),"Error in case 1")

	//case 2
	assert.Equal(t, false, isPalindrome(10),"Error in case 2")

	//case 3
	assert.Equal(t, true, isPalindrome(123454321),"Error in case 3")

	//case 4
	assert.Equal(t, true, isPalindrome(5),"Error in case 4")
}