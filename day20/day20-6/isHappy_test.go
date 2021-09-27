package day20_6

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func isHappy(n int) bool {
	tmp := 0
	num := 0
	hashMap := map[int]struct{}{}
	for{
		for n > 9{
			tmp = n%10
			n = n /10
			num +=tmp*tmp
		}
		num += n*n
		if num == 1{
			return true
		}
		if _,ok := hashMap[num]; ok{
			return false
		}
		hashMap[num] = struct{}{}

		n = num
		num = 0
	}
}

func TestIsHappy(t *testing.T){
	//case 1
	res := isHappy(19)
	assert.Equal(t, true, res, "Error in case 1")
}
