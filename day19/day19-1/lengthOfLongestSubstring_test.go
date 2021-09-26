package day19_1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
//滑动窗口
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {return 0}
	if len(s) == 1{return 1}
	//建立hash
	hashMap := map[uint8]int{}
	count := 1
	left := 0
	right := 0
	hashMap[s[right]] = right

	for {
		//右边界右移
		right++
		//若右边界的值已经在窗口内
		if pos,ok := hashMap[s[right]]; ok{
			//清除不再窗口内的值
			for i := left; i<=pos; i++{
				delete(hashMap, s[i])
			}
			//收拢左边界
			left = pos+1
		}
		//记录右边界的值进入窗口
		hashMap[s[right]] = right
		//记录最大子串长度
		tmp := right-left+1
		if tmp >count{
			count = tmp
		}
		//结束条件
		if right == len(s)-1{
			return count
		}
	}
}

func TestLengthOfLongestSubstring(t *testing.T){
	res := lengthOfLongestSubstring("pwwkew")
	assert.Equal(t, 3, res, "Error in case 1")

}