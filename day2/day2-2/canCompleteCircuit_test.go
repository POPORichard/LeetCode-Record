package day2_2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func canCompleteCircuit(gas []int, cost []int) int {
	length := len(gas)
	var remain []int
	for i:=0;i<length;i++{
		remain = append(remain,gas[i]-cost[i])
	}

	for i:=0;i<length;i++{
		ownOil:= 0
		adder := 0
		success := true
		if remain[i] <0 {continue}
		for t:=0;t<length;t++{
			if i+t == length {adder = 0-i-t}
			ownOil += remain[i+t+adder]
			if ownOil< 0{
				success = false
				break
			}
		}
		if success {
			return i
		}
	}
	return -1
}



func Test_canCompleteCircuit(t *testing.T){
	// case 1
	gas := []int{1,2,3,4,5}
	cost := []int{3,4,5,1,2}
	res := canCompleteCircuit(gas,cost)
	assert.Equal(t,3,res,"Error in case 1")

	// case 2
	gas = []int{2,3,4}
	cost = []int{3,4,3}
	res = canCompleteCircuit(gas,cost)
	assert.Equal(t,-1,res,"Error in case 2")

}