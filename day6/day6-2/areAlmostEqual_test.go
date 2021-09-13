package day6_2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func areAlmostEqual(s1 string, s2 string) bool {
	//依次比较两个字符串相同位置上的字符
	if s1 == s2{return true}
	differ := make([]int,0,len(s2))
	for i,char := range s2{
		if uint8(char) != s1[i]{		//若相同位置的字符不相同则将该位送入数组
			differ = append(differ, i)
		}
	}
	//判断不相同的位置是否只有两个且内容相互一致
	if len(differ) == 2 && s1[differ[0]] == s2[differ[1]] && s1[differ[1]] == s2[differ[0]]{
		return true
	}
	return false


}

func TestAreAlmostEqual(t *testing.T){
	//case 1
	res := areAlmostEqual("bank", "kanb")
	assert.Equal(t, true, res, "Error in case 1")

	//case 2
	res = areAlmostEqual("attack", "defend")
	assert.Equal(t, false, res, "Error in case 2")

	//case 3
	res = areAlmostEqual("abcd", "dcba")
	assert.Equal(t, false, res, "Error in case 3")

	//case 4
	res = areAlmostEqual("siyolsdcjthwsiplccjbuceoxmpjgrauocx","siyolsdcjthwsiplccpbuceoxmjjgrauocx")
	assert.Equal(t, true, res, "Error in case 3")
}