package day15_4_subarraySum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func subarraySum(nums []int, k int) int {
	//子数组个数
	count := 0
	//前缀和
	pre := 0
	//键为前缀和、值为出现次数
	m := map[int]int{}
	m[0] = 1
	//遍历num
	for i := 0; i < len(nums); i++ {
		//更新前缀和
		pre += nums[i]

		//由题：当k == 子数组和 == pre(当前) - pre(过去)即满足条件
		//即 pre(过去) == pre(当前)-k
		//这里查找map中储存的以前的pre出现过的次数，加到count中
		count += m[pre - k]

		//更新map
		m[pre]++
	}
	return count
}

func TestSubarraySum(t *testing.T){
	//case 1
	num := []int{1,2,3}
	res := subarraySum(num, 3)
	assert.Equal(t, 2, res, "Error in case 1")
}



