package day23_6

import (

	"github.com/stretchr/testify/assert"
	"testing"
)

func minEatingSpeed(piles []int, h int) int {
	if len(piles) == 1{
		if piles[0] < h{
			return 1
		}else {
			tmp := piles[0]%h
			if tmp == 0{
				return piles[0]/h
			}else {
				return piles[0]/h+1
			}
		}
	}

	total := 0
	biggest := 0
	for i := range piles{
		total += piles[i]
		if biggest < piles[i]{
			biggest = piles[i]
		}
	}

	min := total/h
	max := biggest
	if min == max {
		return min
	}
	if min == 0{
		min =1
	}

	var check func(speed int)bool
	check = func(speed int)bool {
		cost := 0
		for i := range piles{
			if piles[i] <= speed{
				cost++
			}else {
				tmp := piles[i]
				cost += tmp/speed
				if tmp%speed != 0{
					cost++
				}
			}
			if cost > h{
				return false
			}
		}
		return true
	}

	mid := 0
	for min < max{
		mid = (min + max) / 2
		if check(mid) {
			max = mid
		} else {
			min = mid+1
		}
	}
	return min
}

func TestMinEatingSpeed(t *testing.T){
	//case 1
	piles := []int{30,11,23,4,20}
	res := minEatingSpeed(piles, 5)
	assert.Equal(t, 30, res, "Error in case 1")

	//case 2
	piles = []int{3,6,7,11}
	res = minEatingSpeed(piles, 8)
	assert.Equal(t, 4, res, "Error in case 2")

	//case3
	piles = []int{312884470}
	res = minEatingSpeed(piles, 968709470)
	assert.Equal(t, 1, res, "Error in case 3")
}
