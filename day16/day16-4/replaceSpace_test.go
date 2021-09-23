package day16_4

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
func replaceSpace(s string) string {
	b := []byte(s)
	length := len(b)
	spaceCount := 0
	// 计算空格数量
	for _, v := range b {
		if v == ' ' {
			spaceCount++
		}
	}
	// 扩展原有切片
	tmp := make([]byte, spaceCount*2)
	b = append(b, tmp...)
	i := length - 1
	j := len(b) - 1
	for i >= 0 {
		if b[i] != ' ' {
			b[j] = b[i]
			i--
			j--
		} else {
			b[j] = '0'
			b[j-1] = '2'
			b[j-2] = '%'
			i--
			j = j - 3
		}
	}
	return string(b)
}

func TestReplaceSpace(t *testing.T){
	//case 1
	res := replaceSpace("We are happy.")
	assert.Equal(t, "We%20are%20happy.", res, "Error in case 1")
}
