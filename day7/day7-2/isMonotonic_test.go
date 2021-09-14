package day7_2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func isMonotonic(nums []int) bool {
	for i:=1; i<len(nums); i++{
		if nums[i] == nums[i-1]{
			continue
		}else if nums[i] > nums[i-1]{
			for t:=i+1; t<len(nums); t++{
				if nums[t] < nums[t-1]{return false}
			}
		}else if nums[i] < nums[i-1]{
			for t:=i+1; t<len(nums); t++{
				if nums[t] > nums[t-1]{return false}
			}
		}
		return true
	}
	return true
}

func TestIsMonotonic(t *testing.T){
	//case 1
	nums := []int{1,2,2,3}
	res := isMonotonic(nums)
	assert.Equal(t, true, res, "Error in case 1")

	//case 2
	nums = []int{1,1,1,1}
	res = isMonotonic(nums)
	assert.Equal(t, true, res, "Error in case 2")

	//case 3
	nums = []int{6,5,4,4}
	res = isMonotonic(nums)
	assert.Equal(t, true, res, "Error in case 3")

	//case 4
	nums = []int{1,4,2,5}
	res = isMonotonic(nums)
	assert.Equal(t, false, res, "Error in case 4")

}