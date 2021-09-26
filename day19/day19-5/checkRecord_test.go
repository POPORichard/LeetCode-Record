package day19_5

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func checkRecord(s string) bool {
	return strings.Count(s, "A") < 2 && !strings.Contains(s, "LLL")
}

func TestCheckRecord(t *testing.T){
	//case 1
	res := checkRecord("PPALLL")
	assert.Equal(t, false, res, "Error in case 1")
}
