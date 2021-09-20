package day13_4_toHex

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	//16进制数
	helper := "0123456789abcdef"
	ans := ""

	for num != 0 && len(ans) < 8 {
		//确定16进制最后一位
		temp := num & 15
		ans = string(helper[temp]) + ans

		//移除最后一位
		num >>= 4
	}
	return ans
}

func TestToHex(t *testing.T){
	//case 1
	res := toHex(26)
	assert.Equal(t, "1a", res, "Error in case 1")
}