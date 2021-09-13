package day6_1

import "sort"

func arrayPairSum(nums []int) (ans int) {
	sort.Ints(nums)				//思路一样
	for i := 0; i < len(nums); i += 2 {
		ans += nums[i]
	}
	return
}
