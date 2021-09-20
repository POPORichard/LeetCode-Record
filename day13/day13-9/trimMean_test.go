package day13_9

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func trimMean(arr []int) float64 {
	sort.Ints(arr)
	n := len(arr)
	sum := 0.0
	for i := n / 20; i < n-n/20; i++ {
		sum += float64(arr[i])
	}
	return sum / (float64(n) * 0.9)
}

func TestTrimMean(t *testing.T){
	//case 1
	arr :=[]int{1,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,3}
	res := trimMean(arr)
	assert.Equal(t, 2.0000, res, "Error in case 1 ")
}

