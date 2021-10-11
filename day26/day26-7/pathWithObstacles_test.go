package day26_7

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// arr[i]-i-1 就是这个元素前面有多少个缺失.(如果前面不缺的话 arr[i]应该在索引i-1处)
func findKthPositive(arr []int, k int) int {
	l, r := 0, len(arr)
	for l < r {
		mid := (l + r) >> 1 //(l+r)/2
		if arr[mid]-mid-1 < k {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return k + l
}

func TestFindKthPositive(t *testing.T){
	//case 1
	arr := []int{2,3,4,7,11}
	res := findKthPositive(arr, 5)
	assert.Equal(t, 9, res, "Error in case 1")
}
