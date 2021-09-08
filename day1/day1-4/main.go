package main

import "fmt"

func main(){
	nums := []int{1,5,4,5}
	fmt.Println(maxProduct(nums))
}

func maxProduct(nums []int) int {
	first := 0
	second := 0
	jumpOver := -1
	for i := range nums{
		if nums[i]>first{
			first = nums[i]
			jumpOver = i
		}
	}
	for i := range nums{
		if i == jumpOver{
			continue
		}
		if nums[i]>second{
			second = nums[i]
		}
	}
	return (first-1)*(second-1)
}
