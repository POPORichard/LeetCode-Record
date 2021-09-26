package day19_3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func intersect(nums1 []int, nums2 []int) []int {
	res := make([]int ,0, 10)
	hashMap := map[int]int{}

	for i := range nums1{
		hashMap[nums1[i]]++
	}

	for i := range nums2{
		if v,ok := hashMap[nums2[i]]; ok && v > 0{
			res = append(res, nums2[i])
			hashMap[nums2[i]]--
		}
	}

	return res
}

func TestIntersect(t *testing.T){
	//case 1
	num1 := []int{4,9,5}
	num2 := []int{9,4,9,8,4}
	res := intersect(num1, num2)
	assert.Equal(t, []int{9,4}, res, "Error in case 1")
}
