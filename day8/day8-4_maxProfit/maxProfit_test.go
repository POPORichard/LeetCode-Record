package day8_4_maxProfit

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//遍历
//超时
func maxProfit(prices []int) int {
	gain := 0
	for i := 0; i<len(prices)-1; i++{
		for t := i+1; t<len(prices); t++{
			if prices[t]<prices[i]{continue}
			if prices[t]-prices[i] > gain{
				gain = prices[t]-prices[i]
			}
		}
	}
	return gain
}

func TestMaxProfit(t *testing.T){
	//case 1
	prices := []int{7,1,5,3,6,4}
	res := maxProfit(prices)
	assert.Equal(t, 5, res, "Error in case 1")

}
