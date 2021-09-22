package day15_7_minDifference

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func minDifference(nums []int) int {
	//当长度小于5时可直接返回
	if len(nums) < 5 {return 0}
	//排序从小到大
	sort.Ints(nums)

	//维护一个固定长度的窗口，使窗口外的值的个数为3
	min := nums[len(nums)-1]- nums[3]
	//移动该窗口找最小能达到的值
	for i:=0; i<3; i++{
		tmp := nums[len(nums)-2-i] - nums[2-i]
		if min > tmp{
			min = tmp
		}
	}
	return min
}

func TestMinDifference(t *testing.T){
	//case 1
	nums := []int{6,6,0,1,1,4,6}
	res := minDifference(nums)
	assert.Equal(t, 2, res, "Error in case 1")
}
