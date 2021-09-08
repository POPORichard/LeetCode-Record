package day1_7

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func findPoisonedDuration(timeSeries []int, duration int) int {
	res := 0
	for i:=1;i<len(timeSeries);i++{
		if timeSeries[i]-timeSeries[i-1] < duration{
			res += timeSeries[i]-timeSeries[i-1]
		}else {
			res += duration
		}
	}
	res += duration
	return res
}

func Test_findPoisonedDuration(t *testing.T){
	//case 1
	timeSeries := []int{1,4}
	res := findPoisonedDuration(timeSeries,2)
	assert.Equal(t, 4, res, "Error in case 1")

	//case 2
	timeSeries = []int{1,2}
	res = findPoisonedDuration(timeSeries,2)
	assert.Equal(t, 3, res, "Error in case 1")

}