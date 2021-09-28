package day21_8

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func relativeSortArray(arr1 []int, arr2 []int) []int {
	n := len(arr1)
	k := 0
	for _, num := range arr2 {
		sort.Slice(arr1[k:], func(i, j int) bool {
			return arr1[k+i] == num
		})
		for k < n && arr1[k] == num {
			k++
		}
	}
	sort.Ints(arr1[k:])
	return arr1
}

func testRelativeSortArray(t *testing.T){
	//case 1
	arr1 := []int{2,3,1,3,2,4,6,7,9,2,19}
	arr2 := []int{2,1,4,3,9,6}
	res := relativeSortArray(arr1, arr2)
	assert.Equal(t, []int{2,2,2,1,4,3,3,9,6,7,19}, res, "Error in case 1")
}