package day12_8

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func intersection(nums1 []int, nums2 []int) []int {
	res := make([]int, 0, len(nums2))
	hashMap := map[int]bool{}
	for i := range nums1{
		hashMap[nums1[i]] = true
	}
	for i := range nums2{
		flag := hashMap[nums2[i]]
		if flag {
			delete(hashMap, nums2[i])
			res = append(res, nums2[i])
		}
	}
	return res
}

func TestIntersection(t *testing.T){
	//case 1
	nums1 := []int{4,9,5}
	nums2 := []int{9,4,9,8,4}
	res := intersection(nums1, nums2)
	assert.Equal(t, []int{9, 4}, res , "Error in case 1")
}
