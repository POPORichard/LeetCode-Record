package day23_9

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type data struct {
	num int
	pos int
}

func replaceElements(arr []int) []int {
	stack := make([]data,len(arr))

	for i := range arr{
		for len(stack) != 0{
			if stack[len(stack)-1].num <= arr[i]{
				stack = stack[:len(stack)-1]
				continue
			}
			stack = append(stack, data{
				num: arr[i],
				pos: i,
			})
			break
		}

		if len(stack)==0{
			stack = append(stack, data{
				num: arr[i],
				pos: i,
			})
			continue
		}

	}

	res := make([]int,len(arr))
	t := 0
	for i :=range stack{
		for ; t<stack[i].pos; t++{
			res[t] = stack[i].num
		}
		t = stack[i].pos
	}
	res[len(res)-1] = -1
	return res
}

func TestReplaceElements(t *testing.T){
	arr := []int{17, 18, 5, 4, 6, 1}
	res := replaceElements(arr)
	assert.Equal(t, []int{18,6,6,6,1,-1}, res, "Error in case 1")
}
