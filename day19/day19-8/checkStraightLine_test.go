package day19_8

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func checkStraightLine(coordinates [][]int) bool {
	if len(coordinates) == 2 {return true}
	for i :=1; i < len(coordinates) - 1;i++{
		k1 := (coordinates[i][1] - coordinates[i-1][1]) * (coordinates[i+1][0] - coordinates[i][0])
		k2 := (coordinates[i+1][1] - coordinates[i][1]) * (coordinates[i][0] - coordinates[i-1][0])
		if k1 != k2{
			return false
		}
	}
	return true
}

func TestCheckStraightLine(t *testing.T){
	//case1
	coordinates := [][]int{{1,2},{2,3},{3,4},{4,5},{5,6},{6,7}}
	res := checkStraightLine(coordinates)
	assert.Equal(t, true, res, "Error in case 1")
}
