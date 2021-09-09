package day2_10_findMiddleIndex

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func findMiddleIndex(nums []int) int {
	tot := 0
	for _, v := range nums {	//求数组总和
		tot += v
	}
	s := 0
	//从第一个数值遍历
	//若总和减去该数值刚好等于该数值前面的数值和的2倍，该数值即在中间位置
	for i, v := range nums {
		if s*2 == tot-v {
			return i
		}
		s += v
	}
	return -1
}

func Test_findMiddleIndex(t *testing.T){
	//case 1
	nums := []int{2,3,-1,8,4}
	res := findMiddleIndex(nums)
	assert.Equal(t, 3, res ,"Error in case 1")

	//case 2
	nums = []int{1,-1,4}
	res = findMiddleIndex(nums)
	assert.Equal(t, 2, res ,"Error in case 2")

	//case 3
	nums = []int{2,5}
	res = findMiddleIndex(nums)
	assert.Equal(t, -1, res ,"Error in case 3")

	//case 4
	nums = []int{1,1}
	res = findMiddleIndex(nums)
	assert.Equal(t, -1, res ,"Error in case 4")

	//case 5
	nums = []int{4,0}
	res = findMiddleIndex(nums)
	assert.Equal(t, 0, res ,"Error in case 5")

	//case 6
	nums = []int{1,1,1,1}
	res = findMiddleIndex(nums)
	assert.Equal(t, -1, res ,"Error in case 6")

	//case 7
	nums = []int{0,0,0,0}
	res = findMiddleIndex(nums)
	assert.Equal(t, 0, res ,"Error in case 7")

	//case 8
	nums = []int{3,2,-1,-4,8}
	res = findMiddleIndex(nums)
	assert.Equal(t, 1, res ,"Error in case 8")
}