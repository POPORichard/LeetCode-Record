package day17_2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func countNicePairs(nums []int) int {
	hashMap := map[int]int{}
	count := 0
	for i := range nums{
		t := nums[i]-rev(nums[i])
		if _,ok := hashMap[t]; ok{
			count += hashMap[t]
		}
		hashMap[t]++
	}
	return count%1000000007
}

func rev(num int) int{
	rev := 0
	for num>9{
		rev += num%10
		num = num/10
		rev = rev*10
	}
	rev += num
	return rev
}


func TestCountNicePairs(t *testing.T){
	//case 1
	nums := []int{13, 10, 35, 24, 76}
	res := countNicePairs(nums)
	assert.Equal(t, 4, res, "Error in case 1")
}