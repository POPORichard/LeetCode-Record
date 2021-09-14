package day7_6

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func getConcatenation(nums []int) []int {
	ans := make([]int,0,len(nums)*2)	//提前分配足够的内存
	ans = append(nums,nums...)
	return ans
}

func TestGetConcatenation(t *testing.T){
	//case1
	nums := []int{1,3,2,1}
	res := getConcatenation(nums)
	assert.Equal(t, []int{1,3,2,1,1,3,2,1}, res, "Error in case 1")

}
