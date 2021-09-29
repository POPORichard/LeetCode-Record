package day22_3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func maxScore(cardPoints []int, k int) int {
	cover := len(cardPoints) - k
	res := 0
	for i:= range cardPoints[cover:]{
		res += cardPoints[cover:][i]
	}
	if cover == 0{
		return res
	}

	tmp :=res
	right := cover-1

	for ;right < len(cardPoints)-1; right++{
		tmp += cardPoints[right-cover+1] - cardPoints[right+1]
		if tmp > res{
			res = tmp
		}
	}
	return res
}

func TestMaxScore(t *testing.T){
	//case 1
	cardPoints := []int{1,2,3,4,5,6,1}
	res := maxScore(cardPoints, 3)
	assert.Equal(t, 12, res, "Error in case 1")

	//case 2
	cardPoints = []int{2,2,2}
	res = maxScore(cardPoints, 2)
	assert.Equal(t, 4, res, "Error in case 2")
}
