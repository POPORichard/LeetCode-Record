package day23_3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func minSwaps(s string) int {
	zero, one := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			zero++
		} else {
			one++
		}
	}
	if abs(zero, one) > 1 {
		return -1
	}
	if zero == one {
		return min(find(s, '0'), find(s, '1'))
	} else if zero > one {
		// 0比1多，01010形式才可以，不能以1开头
		return find(s, '0')
	} else {
		// 同理
		return find(s, '1')
	}
}
func abs(a, b int) int {
	if a < b {
		return b - a
	} else {
		return a - b
	}
}
func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
func find(s string, start byte) int{
	diff := 0
	for i := 0; i < len(s); i++ {
		if s[i] != start {
			diff++
		}
		if start == '1' {
			start = '0'
		} else {
			start = '1'
		}
	}
	//fmt.Println(start, " ", diff)
	return diff / 2;
}

func TestMinSwaps(t *testing.T){
	//case 1
	res := minSwaps("111000")
	assert.Equal(t, 1, res, "Error in case 1")

	//case 2
	res = minSwaps("1110")
	assert.Equal(t, -1,res, "Error in case 1")

	//case 3
	res = minSwaps("110")
	assert.Equal(t, 1, res, "Error in case 1")
}
