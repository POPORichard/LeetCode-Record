package day8_9

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
func threeConsecutiveOdds(arr []int) bool {
	t := 0
	for i := range arr{
		if arr[i]%2 == 1{
			t++
			if t == 3{return true}
		}else {
			t = 0
		}
	}
	return false
}

func TestThreeConsecutiveOdds(t *testing.T){
	//case 1
	arr := []int{1,2,34,3,4,5,7,23,12}
	res := threeConsecutiveOdds(arr)
	assert.Equal(t, true, res, "Error in case 1")

}