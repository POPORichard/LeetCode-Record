package day11_9

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
func getWinner(arr []int, k int) int {
	n:=arr[0]
	m:=k
	for i:=1;i<len(arr);i++{
		if arr[i]>n{
			n = arr[i]
			m=k
			m--
		}else {
			m--
		}
		if m==0{
			break
		}
	}
	return n
}

func TestGetWinner(t *testing.T){
	//case 1
	arr := []int{1,9,8,2,3,7,6,4,5}
	res := getWinner(arr, 7)
	assert.Equal(t, 9, res, "Error in case 1")

	//case 2
	arr = []int{1,11,22,33,44,55,66,77,88,99}
	res = getWinner(arr, 1000000000)
	assert.Equal(t, 99, res, "Error in case 2")

	//case 3
	arr = []int{1,25,35,68,42,70}
	res = getWinner(arr, 3)
	assert.Equal(t, 70, res, "Error in case 2")
}

