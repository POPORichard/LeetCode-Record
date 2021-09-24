package day17_1

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func sortArray(nums []int) []int {
	sort.Ints(nums)
	return nums
}

func TestSortArray(t *testing.T){
	//case 1
	nums := []int{5,1,1,2,0,0}
	res := sortArray(nums)
	assert.Equal(t, []int{0,0,1,1,2,5}, res, "Error in case 1")
}