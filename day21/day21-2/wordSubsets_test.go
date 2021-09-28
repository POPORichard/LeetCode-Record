package day21_2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func wordSubsets(words1 []string, words2 []string) []string {
	bMap:=make(map[byte]int)
	for i:=0;i<len(words2);i++{
		temp:=make(map[byte]int)
		for j:=0;j<len(words2[i]);j++{
			temp[words2[i][j]]++
		}
		for key,value:=range temp{
			if bMap[key]<value{
				bMap[key] = value
			}
		}
	}
	res:=make([]string,0)
	for i:=0;i<len(words1);i++{
		aMap:=make(map[byte]int)
		for j:=0;j<len(words1[i]);j++{
			aMap[words1[i][j]]++
		}
		flag:=0
		for key,value:=range bMap{
			if aMap[key]<value{
				flag=1
			}
		}
		if flag == 0{
			res = append(res,words1[i])
		}

	}
	return res
}

func TestWordSubsets(t *testing.T){
	//case 1
	words1 := []string{"amazon","apple","facebook","google","leetcode"}
	words2 := []string{"e", "o"}
	res := wordSubsets(words1, words2)
	assert.Equal(t, []string{"facebook","google","leetcode"}, res, "Error in case 1")
}
