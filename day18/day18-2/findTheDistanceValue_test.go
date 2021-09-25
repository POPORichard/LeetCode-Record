package day18_2

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	cnt := 0
	for _,v1 := range arr1 {
		flag := 1
		for _, v2 := range arr2 {
			if int(math.Abs(float64(v1 - v2))) <= d {
				flag = 0
				break
			}
		}
		if flag == 1 {
			cnt++
		}
	}
	return cnt
}

func TestFindTheDistanceValue(t *testing.T){
	//case 1
	arr1 := []int{4,5,8}
	arr2 := []int{10,9,1,8}
	res := findTheDistanceValue(arr1, arr2, 2)
	assert.Equal(t, 2, res, "Error in case 1")

}