package day9_9

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func numRabbits(answers []int) int {
	hashMap := map[int]int{}
	for i:= range answers{
		hashMap[answers[i]]++
	}
	 num := 0
	for y, x := range hashMap {
		num += (x + y) / (y + 1) * (y + 1)
	}
	 return num
}

func TestNumRabbits(t *testing.T){
	//case 1
	answer := []int{1,1,2}
	res := numRabbits(answer)
	assert.Equal(t, 5, res, "Error in case 1")

	//case 2
	answer = []int{1,0,1,0,0}
	res = numRabbits(answer)
	assert.Equal(t, 5, res, "Error in case 1")
}