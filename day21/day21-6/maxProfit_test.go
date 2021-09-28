package day21_6

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)
//time out
func maxProfit(inventory []int, orders int) int {
	sort.Ints(inventory)
	pointer := len(inventory)-1
	res := 0
	for i :=0; i<orders; i++{
		res += inventory[pointer]
		inventory[pointer]--

		if len(inventory) > 1{
			for{
				if pointer<len(inventory)-1 && inventory[pointer] < inventory[pointer+1]{
					pointer++
				}else if pointer>0 && inventory[pointer] < inventory[pointer-1]{
					pointer--
				}else {
					break
				}
			}
		}

	}
	return res % 1000000007
}

func TestMaxProfit(t *testing.T){
	//case 1
	inventory := []int{2,5}
	res := maxProfit(inventory, 4)
	assert.Equal(t, 14, res, "Error in case 1")

	//case 2
	inventory = []int{3,5}
	res = maxProfit(inventory, 6)
	assert.Equal(t, 19, res, "Error in case 1")
}