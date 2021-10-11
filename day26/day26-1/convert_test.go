package day26_1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func convert(s string, numRows int) string {
	if numRows == 1 || len(s) <= numRows {
		return s
	}
	groupLen := numRows*2 - 2
	var ansString []byte

	for i := 0; i < numRows; i++ {
		//计算第 i 行字符串
		left := i
		right := groupLen - left
		for {
			if left >= len(s) {
				break
			}
			ansString = append(ansString, s[left])
			left += groupLen
			if i != 0 && i != numRows-1 {
				if right >= len(s) {
					break
				}
				ansString = append(ansString, s[right])
				right += groupLen
			}
		}
	}
	return string(ansString)
}

func TestConvert(t *testing.T){
	//case 1
	res := convert("PAYPALISHIRING", 3)
	assert.Equal(t, "PAHNAPLSIIGYIR", res, "Error in case 1")
}
