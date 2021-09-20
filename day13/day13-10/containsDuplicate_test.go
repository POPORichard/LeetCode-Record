package day13_10

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func containsDuplicate(nums []int) bool {
	set := map[int]struct{}{}
	for _, v := range nums {
		if _, has := set[v]; has {
			return true
		}
		set[v] = struct{}{}
	}
	return false
}

func TestContainsDuplicate(t *testing.T){
	nums := []int{1,1,1,3,3,4,3,2,4,2}
	res := containsDuplicate(nums)
	assert.Equal(t, true, res, "Error in case 1")
}