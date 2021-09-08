package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func numEquivDominoPairs(dominoes [][]int) int {
	res :=0
	cnt := [100]int{}
	for _, domino := range dominoes{
		if domino[0] > domino[1]{
			domino[0], domino[1] = domino[1], domino[0]
		}
		pos := domino[0]*10 + domino[1]
		res += cnt[pos]
		cnt[pos]++
	}
	return res
}


func Test_numEquivDominoPairs(t *testing.T) {
	//case 1
	dominoes := [][]int{{1, 2}, {2, 1}, {3, 4}, {5, 6}}
	res := numEquivDominoPairs(dominoes)
	assert.Equal(t, 123, res, "they should be equal")
}