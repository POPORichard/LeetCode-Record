package day11_8

func maxScoreSightseeingPair(values []int) int {
	ans, mx := 0, values[0] + 0
	for i := 1; i < len(values); i++ {
		ans = max(ans, mx + values[i] - i)
		mx = max(mx, values[i] + i)
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
