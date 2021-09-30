package day10_4

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func selfDividingNumbers(left int, right int) []int {
	res := make([]int ,0, right-left)
	for i:=left;i<=right;i++{
		if isDividingNumbers(i){
			res = append(res,i)
		}
	}
	return res
}

func isDividingNumbers(i int)bool{
	if i<10{
		return true
	}
	v:=i
	for i!=0{
		r :=i%10
		if r==0||v%r!=0{
			return false
		}
		i =i/10
	}
	return true
}


func TestSelfDividingNumbers(t *testing.T){
	res := selfDividingNumbers(1, 22)
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22}, res, "Error in case 1")

}


