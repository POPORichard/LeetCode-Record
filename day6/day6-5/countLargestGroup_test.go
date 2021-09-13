package day6_5

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func countLargestGroup(n int) int {
	//建立hash表
	hashMap := map[int]int{}
	for i:=1; i<=n; i++{
		hashMap[addNum(i)]++
	}
	//找值最大的，并计数
	res := 0
	biggest := hashMap[1]
	for i := range hashMap{
		if hashMap[i] == biggest{
			res++
		}else if hashMap[i] > biggest{
			res = 1
			biggest = hashMap[i]
		}
	}
	return res
}
//算各个位上的加和
func addNum(n int)int{
	res := 0
	for{
		if n>9{
			res += n%10
			n = n/10
		}else {
			res += n
			return res
		}
	}
}

func TestCountLargestGroup(t *testing.T){
	//case 1
	res := countLargestGroup(24)
	assert.Equal(t, 5, res, "Error in case 1")

}

