package day16_5

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func getHint(secret string, guess string) string {
	a := 0
	b := 0

	hashMapA := map[byte]int{}
	hashMapB := map[byte]int{}

	for i := range secret{
		if secret[i] == guess[i]{
			a++
			continue
		}
		hashMapA[secret[i]]++
		hashMapB[guess[i]]++
	}
	for byt,n := range hashMapA{
		if t,ok := hashMapB[byt]; ok{
			b += min(n,t)
		}
	}
	return strconv.Itoa(a)+"A"+strconv.Itoa(b)+"B"
}

func min (a,b int)int{
	if a>b{
		return b
	}
	return a
}

func TestGetHint(t *testing.T){
	//case 1
	res := getHint("1807","7810")
	assert.Equal(t, "1A3B", res, "Error in case 1")
}

