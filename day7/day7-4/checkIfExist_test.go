package day7_4

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func checkIfExist(arr []int) bool {
	hash := map[int]int{}
	for i := range arr{
		if arr[i]%2 == 0{
			hash[arr[i]] = i	//构建hash表
		}
	}
	if len(hash) == 0 {return false}
	for i := range arr{
		t,ok := hash[arr[i]*2]
		if ok && t != i {return true}	//0可能会找到到自己，需要判断位置
	}
	return false
}

func TestCheckIfExist(t *testing.T){
	//case 1
	arr := []int{10,2,5,3}
	res := checkIfExist(arr)
	assert.Equal(t, true, res, "Error in case 1")

	//case 2
	arr = []int{3,1,7,11}
	res = checkIfExist(arr)
	assert.Equal(t, false, res, "Error in case 2")

	//case 3
	arr =[]int{-778,-481,842,495,44,1000,-572,977,240,-116,673,997,-958,-539,-964,-187,-701,-928,472,965,-672,-88,443,36,388,-127,115,704,-549,1000,998,291,633,423,57,-77,-543,72,328,-938,-192,382,179}
	res = checkIfExist(arr)
	assert.Equal(t, true, res, "Error in case 3")

	//case 4
	arr = []int{-2,0,10,-19,4,6,-8}
	res = checkIfExist(arr)
	assert.Equal(t, false, res, "Error in case 4")


}