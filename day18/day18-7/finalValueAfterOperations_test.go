package day18_7

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func finalValueAfterOperations(operations []string) (ans int) {
	for _, op := range operations {
		if op[1] == '+' {
			ans++
		} else {
			ans--
		}
	}
	return
}

func TestFinalValueAfterOperations(t *testing.T){
	//case 1
	res := finalValueAfterOperations([]string{"--X","X++","X++"})
	assert.Equal(t, 1, res, "Error in case 1")

}

