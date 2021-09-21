package day14_7

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func maximumTime(time string) string {
	t := []byte(time)
	if t[0] == '?' {
		if '4' <= t[1] && t[1] <= '9' {
			t[0] = '1'
		} else {
			t[0] = '2'
		}
	}
	if t[1] == '?' {
		if t[0] == '2' {
			t[1] = '3'
		} else {
			t[1] = '9'
		}
	}
	if t[3] == '?' {
		t[3] = '5'
	}
	if t[4] == '?' {
		t[4] = '9'
	}
	return string(t)
}

func TestMaximumTime(t *testing.T){
	//case 1
	res := maximumTime("2?:?0")
	assert.Equal(t, "23:50", res, "Error in case 1")
}