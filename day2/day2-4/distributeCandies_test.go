package day2_4

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func distributeCandies(candies int, num_people int) []int {
	res := make([]int,num_people,num_people)
	i:=0
	pos := 0
	for {
		cost := pos+1+num_people*i
		candies = candies-cost
		if candies >= 0{
			res[pos] +=cost
		}else {
			res[pos] += candies+cost
			break
		}
		pos++
		if pos >= num_people{
			i++
			pos = 0
		}
	}
	return res
}

func Test_distributeCandies(t *testing.T){

	//case 1
	res := distributeCandies(7,4)
	assert.Equal(t,[]int{1,2,3,1},res,"Error in case 1")

	//case 2
	res = distributeCandies(10,3)
	assert.Equal(t,[]int{5,2,3},res,"Error in case 1")

}