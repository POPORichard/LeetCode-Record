package day13_3

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func searchInsert(nums []int, target int) int {
	return sort.SearchInts(nums, target)
}

func TestSearchInsert(t *testing.T){
	nums := []int{1,3,5,6}
	res := searchInsert(nums, 5)
	assert.Equal(t, 2, res, "Error in case 1")
}