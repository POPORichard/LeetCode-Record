package day12_5

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func sumOfUnique(nums []int) int {
	ans := 0
	//建立哈希
	maps := make(map[int]int)
	for _, v := range nums {
		maps[v]++
	}
	for k, v := range maps {
		if v == 1 {
			ans += k
		}
	}
	return ans
}

func TestSumOfUnique(t *testing.T){
	//case 1
	nums := []int{1,2,3,4,5}
	res := sumOfUnique(nums)
	assert.Equal(t, 15, res, "Error in case 1")
}
