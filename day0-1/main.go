package main

import "fmt"

func main(){
	nums := []int{2,1}
	fmt.Println(findKthLargest(nums,2))
}
func findKthLargest(nums []int, k int) int {
	length := len(nums)
	if length == 1{return nums[0]}
	biggest := 0
	tmp := 0
	target := 0
	for i := 0; i<k; i++{
		target = 0
		biggest = nums[i]
		target = 0
		for t := i+1; t<length; t++{
			if biggest<nums[t]{
				biggest = nums[t]
				target = t
			}
		}
		//if i>0{
		//	if biggest == nums[i-1]{
		//		k++
		//	}
		//}
		if target != 0{
			tmp = nums[i]
			nums[i] = nums[target]
			nums[target] = tmp
		}
	}
	return biggest
}



