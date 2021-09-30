package day23_1

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func findLHS(nums []int) int {
	sort.Ints(nums)
	res := 0
	start := 0

	for end := 0 ; end < len(nums) ; end++{
		if nums[end] - nums[start] > 1{
			start++
		}
		if nums[end] - nums[start] == 1{
			tmp :=end-start+1
			if  tmp >res{
				res = tmp
			}
		}
	}
	return res
}

func TestFindLHS(t *testing.T){
	//case 1
	nums := []int{1,3,2,2,5,2,3,7}
	res := findLHS(nums)
	assert.Equal(t, 5, res, "Error in case 1")
}

