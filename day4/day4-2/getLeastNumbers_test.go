package day4_2

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func getLeastNumbers(arr []int, k int) []int {
	for i:=0; i<k; i++{
		least := arr[i]
		target := -1
		for t:=i+1; t<len(arr); t++{
			if arr[t]< least{
				least = arr[t]
				target = t
			}
		}
		if target != -1{
			tmp := arr[i]
			arr[i] = arr[target]
			arr[target] = tmp
		}
	}
	return arr[:k]
}

func TestGetLeastNumbers(t *testing.T){
	//case 1
	arr := []int{3,2,1}
	res := getLeastNumbers(arr,2)
	assert.Equal(t, []int{1,2}, res, "Error in case 1")

}