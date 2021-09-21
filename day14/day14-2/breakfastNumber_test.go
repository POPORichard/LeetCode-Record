package day14_2

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func breakfastNumber(staple []int, drinks []int, x int) int {
	count := 0
	sort.Ints(staple)
	sort.Ints(drinks)

	for i := 0; i < len(staple); i++ {
		if staple[i] > x {
			break
		}
		low := 0
		high := len(drinks)
		 target := x - staple[i]
		//采用二分查找
		for j := len(drinks) - 1; j >= 0 && low < high; j-- {
			var mid = (low + high) / 2
			//比目标值大，说明目标值在当前mid值左边，接着往左分
			if drinks[mid] > target {
				high = mid
			} else {
				//比目标值小，往右+1查找
				low = mid + 1
			}
		}
		//找到小于等于目标值(target)的元素下标，即[0，low]中的值都满足。
		count += low
	}
	return count % 1000000007
}

func TestBreakfastNumber(t *testing.T){
	//case 1
	staple := []int{10,20,5}
	drinks := []int{5,5,2}
	res := breakfastNumber(staple, drinks, 15)
	assert.Equal(t, 6, res, "Error in case 1")
}