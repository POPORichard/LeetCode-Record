package day26_9

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func threeSum(nums []int) [][]int {
	res := make([][]int, 0, 10)

	sort.Ints(nums)

	n := len(nums)
	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if nums[i] > 0 {
			break
		}
		target := -nums[i]
		j := i + 1
		k := n - 1
		//两数之和
		for j < k {
			if j > i+1 && nums[j] == nums[j-1] {
				j++
				continue
			}
			if nums[i]+nums[j] > 0 {
				break
			}
			if nums[j]+nums[k] > target {
				k--
			} else if nums[j]+nums[k] < target {
				j++
			} else {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				j++
			}
		}
	}
	return res

}

func TestThreeSum(t *testing.T) {
	//case 1
	nums := []int{-1, 0, 1, 2, -1, -4}
	res := threeSum(nums)
	assert.Equal(t, [][]int{{-1, -1, 2}, {-1, 0, 1}}, res, "Error in case 1")

}
