package day4_4

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//位运算
func anotherFindErrorNums(nums []int) []int {
	xor := 0
	for _, v := range nums {
		xor ^= v
	}
	n := len(nums)
	for i := 1; i <= n; i++ {
		xor ^= i
	}
	lowbit := xor & -xor
	num1, num2 := 0, 0
	for _, v := range nums {
		if v&lowbit == 0 {
			num1 ^= v
		} else {
			num2 ^= v
		}
	}
	for i := 1; i <= n; i++ {
		if i&lowbit == 0 {
			num1 ^= i
		} else {
			num2 ^= i
		}
	}
	for _, v := range nums {
		if v == num1 {
			return []int{num1, num2}
		}
	}
	return []int{num2, num1}
}

func TestAnotherFindErrorNums(t *testing.T){

	//case 1
	nums := []int{1,2,2,4}
	res := anotherFindErrorNums(nums)
	assert.Equal(t, []int{2,3}, res, "Error in case 1")
}