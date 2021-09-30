package day23_5

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func numTrees(n int) int {
	C := 1
	for i := 0; i < n; i++ {
		C = C * 2 * (2 * i + 1) / (i + 2);
	}
	return C
}

func TestNumTrees(t *testing.T){
	//case 1
	res := numTrees(3)
	assert.Equal(t, 5, res, "Error in case 1")

}
