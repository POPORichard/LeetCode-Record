package day12_3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func numJewelsInStones(jewels string, stones string) int {
	jewelsCount := 0
	jewelsSet := map[byte]bool{}
	for i := range jewels {
		jewelsSet[jewels[i]] = true
	}
	for i := range stones {
		if jewelsSet[stones[i]] {
			jewelsCount++
		}
	}
	return jewelsCount
}

func TestNumJewelsInStones(t *testing.T){
	//case 1
	res := numJewelsInStones("aA", "aAAbbbb")
	assert.Equal(t, 3, res, "Error in case 1")

}