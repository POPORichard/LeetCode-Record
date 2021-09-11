package day4_1

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func getMinDistance(nums []int, target int, start int) int {
	if nums[start] == target {return 0}
	left := 1
	right := 1
	for {
		if	start - left >= 0 {
			if nums[start-left] == target {return left}
			left ++
		}
		if start + right < len(nums){
			if nums[start+right] == target {return right}
			right ++
		}
	}
}

func TestGetMinDistance (t *testing.T){
	//case 1
	nums := []int{1,2,3,4,5}
	res := getMinDistance(nums, 5, 3)
	assert.Equal(t, 1, res, "Error in case 1")

	//case 2
	nums = []int{1,1,1,1,1,1,1,1,1,1}
	res = getMinDistance(nums, 1, 0)
	assert.Equal(t, 0, res, "Error in case 1")

	//case 3
	nums = []int{5,3,6}
	res = getMinDistance(nums, 5, 2)
	assert.Equal(t, 2, res, "Error in case 1")


}