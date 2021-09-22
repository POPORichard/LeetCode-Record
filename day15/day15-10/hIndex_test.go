package day15_10

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func hIndex(citations []int) int {
	sort.Ints(citations)
	times := make([]int,citations[len(citations)-1]+1)

	for i := range citations{
		quote := citations[i]
		for t := 0; t<=quote; t++{
			times[t]++
		}
	}

	ans := 0
	for i := 0 ;i<len(times); i++{
		if times[i] < i{
			return ans
		}
		ans = i
	}
	return ans
}

func TestHIndex(t *testing.T){
	//case 1
	citations := []int{3,0,6,1,5}
	res := hIndex(citations)
	assert.Equal(t, 3, res, "Error in case 1")

	//case 2
	citations = []int{1}
	res = hIndex(citations)
	assert.Equal(t, 1, res, "Error in case 2")

}
