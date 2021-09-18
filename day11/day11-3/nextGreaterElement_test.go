package day11_3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	hashMap := map[int]int{}
	for i := range nums2{
		bigger := -1
		for t:=i+1; t<len(nums2); t++{
			if nums2[t] > nums2[i]{
				bigger = nums2[t]
				break
			}
		}
		hashMap[nums2[i]] = bigger
	}

	ans := make([]int, 0, len(nums1))
	for i := range nums1{
		ans = append(ans, hashMap[nums1[i]])
	}

	return ans
}

func TestNextGreaterElement(t *testing.T){
	//case 1
	nums1 := []int{4,1,2}
	nums2 := []int{1,3,4,2}
	res := nextGreaterElement(nums1, nums2)
	assert.Equal(t, []int{-1,3,-1}, res, "Error in case 1")

}
