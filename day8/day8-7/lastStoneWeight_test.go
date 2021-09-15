package day8_7

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func lastStoneWeight(stones []int) int {
	length := len(stones)
	if length == 1 {return stones[0]}
	//排序
	for i := 0; i < length; i++ {
		biggest := stones[i]
		target := -1
		for t := i + 1; t < length; t++ {
			if stones[t] > biggest {
				biggest = stones[t]
				target = t
			}
		}
		if target != -1 {
			tmp := stones[i]
			stones[i] = stones[target]
			stones[target] = tmp
		}
	}
	if len(stones) == 2 {return stones[0]-stones[1]}
	//按要求执行
	for {
		left := stones[0] - stones[1]
		stones = stones[2:]

		if left != 0 {
			if len(stones) == 0 {return left}
			if left <= stones[len(stones)-1] {
				stones = append(stones, left)
			} else {
				for t := len(stones) - 1; t >= 0; t-- {
					if left < stones[t] {
						last := make([]int,0,len(stones[t+1:]))
						for i := range stones[t+1:]{
							last = append(last,stones[t+1:][i])	//不能直接last:=stones[t+1:],last使用地址会导致其值改变
						}
						stones = append(stones[:t+1],left)
						stones = append(stones,last...)
						break
					}
					if t == 0{
						stones = append([]int{left},stones...)
					}
				}
			}
		}

		if len(stones) == 1 {
			return stones[0]
		}
		if len(stones) == 0{
			return 0
		}
	}
}

func TestLastStoneWeight(t *testing.T) {
	//case 1
	stones := []int{2, 7, 4, 1, 8, 1}
	res := lastStoneWeight(stones)
	assert.Equal(t, 1, res, "Error in case 1")

	//case 2
	stones = []int{1, 3}
	res = lastStoneWeight(stones)
	assert.Equal(t, 2, res, "Error in case 2")

	//case 3
	stones = []int{3, 7, 8}
	res = lastStoneWeight(stones)
	assert.Equal(t, 2, res, "Error in case 3")

	//case 4
	stones = []int{9, 3, 2, 10}
	res = lastStoneWeight(stones)
	assert.Equal(t, 0, res, "Error in case 4")

	//case 5
	stones = []int{7, 6, 7, 6, 9}
	res = lastStoneWeight(stones)
	assert.Equal(t, 3, res, "Error in case 5")

	//case 6
	stones = []int{2, 6, 6, 9, 4, 3}
	res = lastStoneWeight(stones)
	assert.Equal(t, 0, res, "Error in case 6")

	//case 7
	stones = []int{5, 1, 8, 10, 7}
	res = lastStoneWeight(stones)
	assert.Equal(t, 1, res, "Error in case 7")

	//case 8
	stones = []int{9, 10, 4, 5, 7, 1}
	res = lastStoneWeight(stones)
	assert.Equal(t, 0, res, "Error in case 7")
}
