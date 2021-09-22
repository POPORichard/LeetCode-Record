package day15_8_findMaxLength

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
//将0看作-1
//即找和为0的最长子字符窜
func betterFindMaxLength(nums []int) (maxLength int) {
	mp := map[int]int{0: -1}
	counter := 0
	for i, num := range nums {
		if num == 1 {
			counter++
		} else {
			counter--
		}
		if prevIndex, has := mp[counter]; has {
			maxLength = max(maxLength, i-prevIndex)
		} else {
			mp[counter] = i
		}
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func TestBetterFindMaxLength(t *testing.T){
	//case 1
	nums := []int{1,1,1,0,0,0,1,1,0}
	res := betterFindMaxLength(nums)
	assert.Equal(t, 8, res, "Error in case 1")
}
