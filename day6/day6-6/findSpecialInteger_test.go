package day6_6

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func findSpecialInteger(arr []int) int {
	need := len(arr)/4+1
	tot := 0
	target := arr[0]
	for i := range arr{
		if arr[i] == target{
			tot++
			if tot >= need{return target}
		}else {
			target = arr[i]
			tot = 1
		}
	}
	return -1
}

func TestFindSpecialInteger(t *testing.T){
	//case 1
	arr := []int{1,2,2,6,6,6,6,7,10}
	res := findSpecialInteger(arr)
	assert.Equal(t, 6, res, "Error in case 1")

}
