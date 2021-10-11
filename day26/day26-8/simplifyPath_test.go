package day26_8

import (
	"github.com/stretchr/testify/assert"
	"path"
	"testing"
)

//请务必熟悉标准库
func simplifyPath(p string) string {
	return path.Clean(p)	//看看源码
}

func TestSimplifyPath(t *testing.T){
	//case 1
	res := simplifyPath("/a/./b/../../c/")
	assert.Equal(t, "/c", res, "Error in case 1")
}


