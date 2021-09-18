package day11_10_dailyTemperatures

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type weather struct {
	temp int
	pos  int
}

func dailyTemperatures(temperatures []int) []int {
	length := len(temperatures)
	//建立单调栈
	stack := make([]weather, 0, length)
	ans := make([]int, length)
	//首位数据入栈
	stack = append(stack, weather{temperatures[0], 0})
	for i := 1; i < length; i++ {
		for {
			//若栈顶不为空且栈顶温度小于当前值
			if len(stack) > 0 && stack[len(stack)-1].temp < temperatures[i] {
				//pop
				out := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				//计算天数
				ans[out.pos] = i-out.pos
				//继续判断下一个栈顶
				continue
			}

			//将当前值添加到栈顶
			stack = append(stack, weather{temperatures[i], i})
			break
		}
	}

	return ans
}

func TestDailyTemperatures(t *testing.T) {
	//case 1
	temp := []int{73, 74, 75, 71, 69, 72, 76, 73}
	res := dailyTemperatures(temp)
	assert.Equal(t, []int{1, 1, 4, 2, 1, 1, 0, 0}, res, "Error in case 1")

	//case 2
	temp = []int{30,40,50,60}
	res = dailyTemperatures(temp)
	assert.Equal(t, []int{1, 1, 1, 0}, res, "Error in case 2")
}
