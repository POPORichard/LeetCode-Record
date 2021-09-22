package day15_9

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func maxmiumScore(cards []int, cnt int) int {
	ans := 0
	sort.Sort(sort.Reverse(sort.IntSlice(cards)))
	sum := 0
	for _, v := range cards[:cnt] {
		sum += v
	}
	if sum&1 == 0 {
		return sum
	}
	// 在 cards[cnt:] 中找一个最大的且奇偶性和 x 不同的元素，替换 x
	replace := func(x int) {
		for _, v := range cards[cnt:] {
			if v&1 != x&1 {
				ans = max(ans, sum-x+v)
				break
			}
		}
	}
	replace(cards[cnt-1]) // 替换 cards[cnt-1]
	for i := cnt - 2; i >= 0; i-- {
		if cards[i]&1 != cards[cnt-1]&1 { // 找一个最小的且奇偶性不同于 cards[cnt-1] 的元素，将其替换掉
			replace(cards[i])
			break
		}
	}
	return ans
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func TestMaxmiumScore(t *testing.T){
	//case 1
	cards := []int{1,2,8,9}
	res := maxmiumScore(cards, 3)
	assert.Equal(t, 18, res, "Error in case 1")

	//case 2
	cards = []int{7,4,1}
	res = maxmiumScore(cards, 1)
	assert.Equal(t, 4, res, "Error in case 1")
}
