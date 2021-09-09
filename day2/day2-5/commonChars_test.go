package day2_5

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
//哈希
func commonChars(words []string) []string {
	length := len(words)
	hash := make([]int,length*26,length*26)	//建立哈希表、每个单词占26位
	var res []string
	for i,word := range words{
		for _,char :=range word{
			hash[(i*26)+int(char-97)]++		//将出现过的字母放入哈希
		}
	}

loop:
	for i:=0;i<26;i++{
		if hash[i] != 0{
			less := hash[i]
			for t:=1;t<length;t++{
				if hash[t*26+i] == 0{
					continue loop		//若有一个单词未出现该字母即跳出
				}
				if hash[t*26+i]<less{
					less = hash[t*26+i]	//找最少的出现次数
				}
			}
			for x:=0;x<less;x++{
				res = append(res,string(uint8(97+i)))
			}
		}
	}
	return res
}

func Test_commonChars(t *testing.T){
	//case 1
	words:= []string{"bella", "label", "roller"}
	res := commonChars(words)
	assert.Equal(t,[]string{"e","l","l"},res,"Error in case 1")

	//case 2
	words= []string{"cool","lock","cook"}
	res = commonChars(words)
	assert.Equal(t,[]string{"c","o"},res,"Error in case 2")
}
