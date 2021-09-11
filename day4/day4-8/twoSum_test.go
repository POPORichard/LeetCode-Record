package day4_8

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func twoSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, num := range nums {
		if p, ok := hashTable[target-num]; ok {
			return []int{p, i}
		}
		hashTable[num] = i
	}
	return nil
}

func TestTwoSum(t *testing.T){
	//case 1
	nums := []int{2,7,11,15}
	res := twoSum(nums, 9)
	assert.Equal(t, []int{0,1}, res, "Error in case 1")

}