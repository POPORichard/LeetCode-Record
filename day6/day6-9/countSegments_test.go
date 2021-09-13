package day6_9

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func countSegments(s string) int {
	return len(strings.Fields(s))
}

func TestCountSegments(t *testing.T){
	//case 1
	res := countSegments("Hello, my name is John")
	assert.Equal(t, 5, res, "Error in case 1")

}