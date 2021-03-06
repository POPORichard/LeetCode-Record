package day21_6

func betterMaxProfit(inventory []int, orders int) int {
	// 取模
	var mod int64 = 1e9 + 7
	theMax := 0
	theMin := 0
	for _, v := range inventory {
		if v > theMax {
			theMax = v
		}
	}
	// 找出一个最小的T值，所有值减T之和小于orders
	T := 0
	for {
		// 二分法
		T = (theMax + theMin) / 2

		sum := 0
		largeT := 0
		for _, v := range inventory {
			if v >= T {
				largeT++
				sum += (v - T)
			}
		}
		if sum > orders { // T过小
			theMin = T + 1
		} else if sum+largeT <= orders { // T过大
			theMax = T - 1
		} else {
			break
		}
	}
	var result int64
	r := 0
	for _, v := range inventory {
		if v > T {
			result += (int64)(gaosi(T+1, v)) % mod
			r += v - T
		}
	}

	res := (int64)((orders - r) * T)
	return (int)((result + res) % mod)
}

// 高斯方法求和
func gaosi(a int, b int) int {
	return (int)((a + b) * (b - a + 1) / 2)
}
//
//作者：OneKing
//链接：https://leetcode-cn.com/problems/sell-diminishing-valued-colored-balls/solution/goyu-yan-shuang-bai-zero-de-si-lu-by-oneking/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。