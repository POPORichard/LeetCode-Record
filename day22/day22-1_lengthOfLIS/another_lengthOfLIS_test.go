package day22_1_lengthOfLIS

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func anotherLengthOfLIS(nums []int) int {
	length := len(nums)
	top := make([]int, length)

	piles := 0

	for _, poker := range nums {
		left := 0
		right := piles
		for left < right {
			mid := (left + right) / 2
			if top[mid] >= poker {
				right = mid
			} else {
				left = mid + 1
			}

		}
		if left == piles {
			piles++
		}
		top[left] = poker
	}
	return piles
}

func TestAnotherLengthOfLIS(t *testing.T){
	//case 1
	nums := []int{10,9,2,5,3,7,101,18}
	res := anotherLengthOfLIS(nums)
	assert.Equal(t, 4, res, "Error in case 1")
}
