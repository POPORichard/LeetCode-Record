package day24_8

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func canBeEqual(target []int, arr []int)  bool{
	if len(target) != len(arr) {
		return false
	}
	sort.Ints(target)
	sort.Ints(arr)
	for i := 0; i < len(target); i++ {
		if target[i] != arr[i] {
			return false
		}
	}
	return true
}

func TestCanBeEqual(t *testing.T){
	//case 1
	target := []int{1,2,3,4}
	arr := []int{2,4,1,3}
	res := canBeEqual(target, arr)
	assert.Equal(t, true, res, "Error in case 1")

}
