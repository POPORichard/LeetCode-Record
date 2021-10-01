package day24_3

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func reorderSpaces(text string) string {
	cnt := strings.Count(text, " ")
	fields := strings.Fields(text)
	n := len(fields)
	if n > 1 {
		return strings.Join(fields, strings.Repeat(" ", cnt/(n-1))) + strings.Repeat(" ", cnt % (n-1))
	}
	return strings.Join(fields, "") + strings.Repeat(" ", cnt)
}

func TestReorderSpaces(t *testing.T){
	res := reorderSpaces("  this   is  a sentence ")
	assert.Equal(t, "this   is   a   sentence", res, "Error in case 1")

}
