package day1_10

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func averageWaitingTime(customers [][]int) float64 {
	finishTime := 0
	waitingTime := 0
	for i :=range customers{
		if customers[i][0] > finishTime{
			waitingTime += customers[i][1]
			finishTime = customers[i][0]+customers[i][1]
		}else{
			waitingTime += customers[i][1]+finishTime-customers[i][0]
			finishTime += customers[i][1]
		}
	}
	return float64(waitingTime)/float64(len(customers))
}

func Test_averageWaitingTime(t *testing.T){
	//case 1
	customers := [][]int{{1,2},{2,5},{4,3}}
	res := averageWaitingTime(customers)
	assert.Equal(t,float64(5),res,"Error in case 1")

	//case 2
	customers = [][]int{{5,2},{5,4},{10,3},{20,1}}
	res = averageWaitingTime(customers)
	assert.Equal(t,3.25,res,"Error in case 1")

	//case 3
	customers = [][]int{{2,3},{6,3},{7,5},{11,3},{15,2},{18,1}}
	res = averageWaitingTime(customers)
	assert.Equal(t,4.166666666666667,res,"Error in case 3")

}

