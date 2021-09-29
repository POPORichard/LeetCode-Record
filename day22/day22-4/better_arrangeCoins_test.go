package day22_4

//数学法 n*(n+1)/2 此时的n是行数,方程式的结果是硬币数
func betterArrangeCoins(n int) int {
	l, r := 0, n/2+1
	for l < r {
		//为什么+1,因为下面l=mid.如果剩余两个l将不会缩小范围
		//l=mid,+1
		//r=mid,不+1
		mid := l + (r-l+1)>>1
		val := mid * (mid + 1) / 2
		if val > n {
			r = mid - 1
		} else {
			l = mid
		}
	}
	return l
}
