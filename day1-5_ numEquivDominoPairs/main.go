package main

//方向性错误,不应使用暴力解法
//往哈希方向靠

import "fmt"

func main() {
	dominoes := [][]int{{1, 2}, {2, 1}, {3, 4}, {5, 6}}
	fmt.Println(numEquivDominoPairs(dominoes))
}

func numEquivDominoPairs(dominoes [][]int) int {
	res :=0
	cnt := [100]int{}
	for _, domino := range dominoes{
		if domino[0] > domino[1]{
			domino[0], domino[1] = domino[1], domino[0]
		}
		pos := domino[0]*10 + domino[1]
		res += cnt[pos]
		cnt[pos]++
	}
return res
}
