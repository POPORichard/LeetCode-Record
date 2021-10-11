package day26_4_findMinArrowShots

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	count := 1
	end := points[0][1]

	for i := range points{
		if points[i][0] > end{
			count++
			end = points[i][1]
		}
	}
	return count
}

func TestFindMinArrowShots(t *testing.T){
	points := [][]int{{1,2},{3,4},{5,6},{7,8}}
	res := findMinArrowShots(points)
	assert.Equal(t, 4, res, "Error in case 1")
}
