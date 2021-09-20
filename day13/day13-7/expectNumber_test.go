package day13_7

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func expectNumber(scores []int) int {
	if len(scores) == 0 {
		return 0
	}
	maps := make(map[int]struct{}, 0)
	for _, v := range scores {
		maps[v] = struct{}{}
	}
	return len(maps)
}

func TestExpectNumber(t *testing.T){
	scores := []int{1,2,3}
	res := expectNumber(scores)
	assert.Equal(t, 3, res, "Error in case 1")
}
