package day26_10

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 终点是rounds[len(rounds)-1], 起点是rounds[0], 按题意应该是不中断绕圈,
//那通过次数最多的应该就是最后一圈所经过的扇区了, 所以从终点逆推到起点就是答案

func mostVisited(n int, rounds []int) []int {
	var ret []int

	i, j := rounds[0], rounds[len(rounds)-1]
	for i != j {
		ret = append(ret, j)

		if j == 1 {
			j = n
		} else {
			j--
		}
	}

	ret = append(ret, j)

	res := make([]int,len(ret))
	for i := range ret{
		res[len(res)-1-i] = ret[i]
	}

	return res
}

func TestMostVisited(t *testing.T) {
	//case 1
	rounds := []int{1, 3, 1, 2}
	res := mostVisited(4, rounds)
	assert.Equal(t, []int{1, 2}, res, "Error in case 1")
}
