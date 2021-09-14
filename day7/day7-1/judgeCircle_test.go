package day7_1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
func judgeCircle(moves string) bool {
	pos := []int{0,0}
	for i := range moves{
		switch moves[i] {
		case 'R':
			pos[0]++
		case 'L':
			pos[0]--
		case 'U':
			pos[1]++
		case 'D':
			pos[1]--
		}
	}
	return pos[0] == 0 && pos[1] == 0
}

func TestJudgeCircle(t *testing.T){
	//case 1
	moves := "UD"
	res := judgeCircle(moves)
	assert.Equal(t, true, res, "Error in case 1")
}