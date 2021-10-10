package day25_3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	North = iota
	East
	South
	West
	)

func isRobotBounded(instructions string) bool {
	pos := [2]int{0,0}
	head := North
	for _,i := range instructions{
		switch i {
		case 'L':
			if head == North{
				head = West
			}else {
				head--
			}
		case 'R':
			if head == West{
				head = North
			}else {
				head++
			}
		case 'G':
			switch head {
			case North:
				pos[0]++
			case East:
				pos[1]++
			case South:
				pos[0]--
			case West:
				pos[1]--
			}
		}
	}
	//一次指令之后，只有(x,y)不是原点，并且方向和原来的方向一致，最后才回不去
	if pos != [2]int{0,0} && head == North{
		return false
	}
	return true
}

func TestIsRobotBounded(t *testing.T){
	//case 1
	res := isRobotBounded("GGLLGG")
	assert.Equal(t, true, res, "Error in case 1")
}