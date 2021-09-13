package day6_4

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func numUniqueEmails(emails []string) int {
	//map中key不存在就是增加，存在就是更新
	//利用空map使地址唯一
	m:=map[string]struct{}{}
	for _,email := range emails {
		// 找到"+"的位置
		i:= strings.Index(email, "+")
		// 找到"@"的位置
		j:= strings.LastIndex(email, "@")
		// 如果"+"不存在，前半部分字符串就是j之前的
		if i == -1 {
			i=j
		}
		// 前半部分字符串移除"."，然后和"@"及之后的字符串拼接
		tmp:= strings.ReplaceAll(email[:i],".","") + email[j:]
		m[tmp]=struct{}{}
	}
	return len(m)
}

func TestNumUniqueEmails(t *testing.T){
	//case 1
	emails := []string{"test.email+alex@leetcode.com","test.e.mail+bob.cathy@leetcode.com","testemail+david@lee.tcode.com"}
	res := numUniqueEmails(emails)
	assert.Equal(t, 2, res, "Error in case 1")
}
