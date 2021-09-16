package day9_7_maxWidthRamp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//按题要求使j-i最大，即找到相距最远的满足条件的左指针和右指针
//先以nums[0]为栈底入栈，顺序遍历nums，每次入栈更小的值，建立递减栈
//该递减栈即左指针，且栈顶所对应的值最小
//在最右侧建立右指针
//依次出栈判断左右指针是否满足条件，若满足条件继续出栈以扩大res
//若不满足条件则右指针左移，再次比较，成功则继续出栈继续扩大res
//若右指针值已经小于res，则没有更大的res出现，返回
func maxWidthRamp(nums []int) int {
	//建立单调栈
	stack := make([]int,0,len(nums))
	length := len(nums)
	//从nums[0]开始，将nums中的首个递减序列放入栈中
	for i:=0; i<length; i++ {
		if len(stack)==0 || nums[stack[len(stack)-1]]>nums[i] {
			stack = append(stack, i)
		}
	}

	res := 0
	rightPointer := length - 1
	//当右指针还比结果大时 res = rightPointer - leftPointer
	for rightPointer>res {
		//当栈不为空，且栈底所存左指针对应的值小于等于右指针对应的值（即满足题目要求）
		for len(stack)>0 && nums[stack[len(stack)-1]] <= nums[rightPointer] {
			//当该结果比已存结果好时，替换结果
			if rightPointer-stack[len(stack)-1] > res {
				res = rightPointer-stack[len(stack)-1]
			}
			//弹出栈顶
			stack = stack[:len(stack)-1]
		}
		//右指针左移
		rightPointer--
	}
	return res
}

func TestMaxWidthRamp(t *testing.T){
	//case 1
	nums := []int{6,0,8,2,1,5}
	res := maxWidthRamp(nums)
	assert.Equal(t, 4, res, "Error in case 1")

	//case 2
	nums = []int{9,8,1,0,1,9,4,0,4,1}
	res = maxWidthRamp(nums)
	assert.Equal(t, 7, res, "Error in case 2")

}