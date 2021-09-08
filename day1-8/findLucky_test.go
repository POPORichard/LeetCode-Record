package day1_8

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func findLucky(arr []int) int {
	res := 0
	var t [500]int
	for _,num :=range arr{
		t[num]++
	}
	for i := range t{
		if t[i] == i{
			res = i
		}
	}
	if res == 0{res=-1}
	return res
}

func Test_findLucky(t *testing.T){
	//case 1
	arr := []int{2,2,3,4}
	res := findLucky(arr)
	assert.Equal(t, 2, res ,"Error in case 1")

	//case 2
	arr = []int{2,2,2,3,3}
	res = findLucky(arr)
	assert.Equal(t, -1, res ,"Error in case 2")

}