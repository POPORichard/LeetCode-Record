package day17_10

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type food struct {
	number  int
	outDate int
}

func eatenApples(apples []int, days []int) int {
	count := 0
	length := len(apples)
	queue := make([]food, 0, length)

	for i := 0; ; i++ {
		if i < length {
			apple := food{
				number:  apples[i],
				outDate: days[i] + i,
			}

			if len(queue) == 0 {
				queue = append(queue, apple)
			} else {
				if apple.outDate >= queue[len(queue)-1].outDate {
					queue = append(queue, apple)
				} else if apple.outDate < queue[0].outDate {
					queue = append([]food{apple}, queue...)
				} else {
					for t := range queue {
						if queue[t].outDate > apple.outDate {
							tmp := make([]food, 0, t+1)

							for x := range queue[t:] {
								tmp = append(tmp, queue[t:][x])
							}

							queue = queue[:t]
							queue = append(queue, apple)
							queue = append(queue, tmp...)
							break
						}
					}
				}
			}
		}

		if len(queue) != 0 {
			for {
				if queue[0].outDate == i || queue[0].number == 0 {
					queue = queue[1:]
					if len(queue) == 0 {
						break
					}
					continue
				}
				break
			}
		} else if i > length {
			break
		}

		if len(queue) != 0 {
			queue[0].number--
			count++
		}
	}

	return count
}

func TestEatenApples(t *testing.T){
	//case 1
	apples := []int{1,2,3,5,2}
	days := []int{3,2,1,4,2}
	res := eatenApples(apples, days)
	assert.Equal(t, 7, res, "Error in case 1")

	//case 2
	apples = []int{2,1,10}
	days = []int{2,10,1}
	res = eatenApples(apples, days)
	assert.Equal(t, 4, res, "Error in case 1")
}
