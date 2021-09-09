package day2_2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func betterCanCompleteCircuit(gas []int, cost []int) int {
	for i, n := 0, len(gas); i < n; {
		sumOfGas, sumOfCost, cnt := 0, 0, 0
		for cnt < n {
			j := (i + cnt) % n
			sumOfGas += gas[j]
			sumOfCost += cost[j]
			if sumOfCost > sumOfGas {
				break
			}
			cnt++
		}
		if cnt == n {
			return i
		} else {
			i += cnt + 1
		}
	}
	return -1
}

func Test_betterCanCompleteCircuit(t *testing.T){
	// case 1
	gas := []int{1,2,3,4,5}
	cost := []int{3,4,5,1,2}
	res := betterCanCompleteCircuit(gas,cost)
	assert.Equal(t,3,res,"Error in case 1")

	// case 2
	gas = []int{2,3,4}
	cost = []int{3,4,3}
	res = betterCanCompleteCircuit(gas,cost)
	assert.Equal(t,-1,res,"Error in case 2")

}
