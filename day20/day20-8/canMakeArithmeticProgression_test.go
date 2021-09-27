package day20_8

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func canMakeArithmeticProgression(arr []int) bool {
	sort.Ints(arr)

	for i := 1; i<len(arr)-1; i++{
		if arr[i] - arr[i-1] != arr[i+1] - arr[i]{
			return false
		}
	}
	return true
}

func TestCanMakeArithmeticProgression(t *testing.T){
	//case 1
	arr := []int{3,5,1}
	res := canMakeArithmeticProgression(arr)
	assert.Equal(t, true, res, "Error in case 1")
}