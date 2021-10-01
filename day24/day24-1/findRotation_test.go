package day24_1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func findRotation(a, b [][]int) bool {
	for k := 0; k < 4; k++ {
		a = rotate(a)
		if equal(a, b) {
			return true
		}
	}
	return false
}

func rotate(a [][]int) [][]int {
	n := len(a)
	b := make([][]int, n)
	for i := range b {
		b[i] = make([]int, n)
	}
	for i, r := range a {
		for j, v := range r {
			b[j][n-1-i] = v
		}
	}
	return b
}

func equal(a, b [][]int) bool {
	for i, r := range a {
		for j, v := range r {
			if v != b[i][j] {
				return false
			}
		}
	}
	return true
}

func TestFindRotation(t *testing.T){
	mat := [][]int{{0,1},{1,0}}
	target := [][]int{{1,0},{0,1}}
	res := findRotation(mat, target)
	assert.Equal(t, true, res, "Error in case 1")
}