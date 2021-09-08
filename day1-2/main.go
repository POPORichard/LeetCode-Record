package main

import "fmt"

func main()  {
	widths := []int{10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10}
	S := "abcdefghijklmnopqrstuvwxyz"
	fmt.Println(numberOflines(widths,S))
}

func numberOflines(widths []int, s string)[]int{
	long := 0
	line := 1
	for i:=0;i<len(s);i++{
		long = long + widths[(s[i]-97)]
		if long > 100{
			long = 0
			line++
			i--
		}
		fmt.Println(i)
	}
	res := []int{line, long}
	return res
}