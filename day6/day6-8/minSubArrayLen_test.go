package day6_8

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func minSubArrayLen(target int, nums []int) int {
	length := len(nums)
	leftPointer := 0
	rightPointer := length-1
	if getSum(nums,leftPointer,rightPointer)< target {return 0}
	if getSum(nums,leftPointer,rightPointer) == target {return length}

	//缩小范围
	res := 0
	for i := range nums{
		res += nums[i]
		if res >= target{
			rightPointer = i
			break
		}
	}

	ans := length
	//划动窗口
	for {
		if ans > rightPointer - leftPointer {ans = rightPointer - leftPointer+1}

		if rightPointer == length -1 {
			for {
				if getSum(nums, leftPointer, rightPointer) > target {
					leftPointer++
				}else if getSum(nums, leftPointer, rightPointer) < target{
					leftPointer--
					break
				}else {
					break
				}
			}
			if ans > rightPointer - leftPointer {ans = rightPointer - leftPointer+1}
			return ans

		}

		rightPointer++

		for {
			if getSum(nums, leftPointer, rightPointer) > target {
				leftPointer++
			}else if getSum(nums, leftPointer, rightPointer) < target{
				leftPointer--
				break
			}else {
				break
			}
		}
		continue
	}
}

func getSum(nums []int, left int, right int) int{
	res := 0
	for i:= left; i<=right; i++{
		res += nums[i]
	}
	return res
}

func TestMinSubArrayLen(t *testing.T){
	//case 1
	nums := []int{2,3,1,2,4,3}
	res := minSubArrayLen(7, nums)
	assert.Equal(t, 2, res, "Error in case 1")

	//case 2
	nums = []int{1,4,4}
	res = minSubArrayLen(4, nums)
	assert.Equal(t, 1, res, "Error in case 2")

	//case 3
	nums = []int{1,1,1,1,1,1,1,1}
	res = minSubArrayLen(11, nums)
	assert.Equal(t, 0, res, "Error in case 3")

	//case 4
	nums = []int{10,2,3}
	res = minSubArrayLen(6, nums)
	assert.Equal(t, 1, res, "Error in case 4")

	//case 5
	nums = []int{1,2,3,4,5}
	res = minSubArrayLen(11, nums)
	assert.Equal(t, 3, res, "Error in case 5")

	//case 6
	nums = []int{5,1,3,5,10,7,4,9,2,8}
	res = minSubArrayLen(15, nums)
	assert.Equal(t, 2, res, "Error in case 6")


	//case 7
	nums = []int{2,16,14,15}
	res = minSubArrayLen(20, nums)
	assert.Equal(t, 2, res, "Error in case 7")


}