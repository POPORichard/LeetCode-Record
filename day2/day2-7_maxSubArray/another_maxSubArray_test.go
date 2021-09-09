package day2_7_maxSubArray

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//线段树？？？
func anotherMaxSubArray(nums []int) int {
	return get(nums, 0, len(nums) - 1).mSum;
}

func pushUp(l, r Status) Status {
	iSum := l.iSum + r.iSum
	lSum := max(l.lSum, l.iSum + r.lSum)
	rSum := max(r.rSum, r.iSum + l.rSum)
	mSum := max(max(l.mSum, r.mSum), l.rSum + r.lSum)
	return Status{lSum, rSum, mSum, iSum}
}

func get(nums []int, l, r int) Status {
	if l == r {
		return Status{nums[l], nums[l], nums[l], nums[l]}
	}
	m := (l + r) >> 1
	lSub := get(nums, l, m)
	rSub := get(nums, m + 1, r)
	return pushUp(lSub, rSub)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

type Status struct {
	lSum, rSum, mSum, iSum int
}

func Test_anotherMaxSubArray(t *testing.T){
	//case 1
	nums := []int{-2,1,-3,4,-1,2,1,-5,4}
	res := anotherMaxSubArray(nums)
	assert.Equal(t,6,res,"Error in case 1")

	//case 2
	nums = []int{-1,-2}
	res = anotherMaxSubArray(nums)
	assert.Equal(t,-1,res,"Error in case 2")
}