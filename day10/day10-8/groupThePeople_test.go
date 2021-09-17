package day10_8

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func groupThePeople(groupSizes []int) [][]int {
	//建立Hash
	positionHash := map[int][]int{}
	for i:= range groupSizes{
		positionHash[groupSizes[i]] = append(positionHash[groupSizes[i]],i)
	}

	res := make([][]int,0,len(groupSizes))
	for group, pos := range positionHash{
			//切割group
			for len(pos) != 0{
				res = append(res,pos[:group])
				pos = pos[group:]
			}
	}
	return res
}

func TestGroupThePeople(t *testing.T){
	//case 1
	groupSizes := []int{3,3,3,3,3,1,3}
	res := groupThePeople(groupSizes)
	assert.Equal(t, [][]int{{0,1,2},{3,4,6},{5}}, res,"Error in case 1")
}
