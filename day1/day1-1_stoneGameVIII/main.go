package main

import (
	"fmt"
)

//未作出 :没有考虑动态规划

func main(){
	stones := []int{7,-6,5,10,5,-2,-6}
	fmt.Println(stoneGameVIII(stones))
}

func stoneGameVIII(stones []int)int{
	length := len(stones)
	pre := make([]int,0,length)
	score := 0
	for i := range stones{
		score = score + stones[i]
		pre = append(pre, score)
	}

	f := make([]int,length,length)
	f[length-1] = pre[length-1]

	for i:= length-2;i>=1; i--{
		f[i] = max(f[i+1],pre[i]-f[i+1])
	}

	return f[1]
}

func max(a int,b int)int{
	if a > b{
		return a
	}
	return b
}

