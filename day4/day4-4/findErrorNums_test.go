package day4_4

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//哈希
func findErrorNums(nums []int) []int {
	cnt := map[int]int{}
	for _, v := range nums {
		cnt[v]++
	}
	ans := make([]int, 2)
	for i := 1; i <= len(nums); i++ {
		if c := cnt[i]; c == 2 {
			ans[0] = i
		} else if c == 0 {
			ans[1] = i
		}
	}
	return ans
}

func TestFindErrorNums(t *testing.T){

	//case 1
	nums := []int{1,2,2,4}
	res := findErrorNums(nums)
	assert.Equal(t, []int{2,3}, res, "Error in case 1")
}