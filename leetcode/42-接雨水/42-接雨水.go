package main

// 42-接雨水
func trap(height []int) int {
	/*
	   water[i] = min(
	       # 左边最⾼的柱⼦
	       max(height[0..i]),
	       # 右边最⾼的柱⼦
	       max(height[i..end])
	       ) - height[i]
	*/
	if len(height) == 0 {
		return 0
	}
	res := 0
	n := len(height)
	lMax, rMax := make([]int, n), make([]int, n)
	lMax[0] = height[0]
	rMax[n-1] = height[n-1]
	// 从左往右找到最大值
	for i := 1; i < n; i++ {
		lMax[i] = max(height[i], lMax[i-1])
	}
	// 从右往左找到最大值
	for i := n - 2; i >= 0; i-- {
		rMax[i] = max(height[i], rMax[i+1])
	}
	// 累加结果
	for i := 1; i < n-1; i++ {
		res += (min(lMax[i], rMax[i]) - height[i])
	}
	return res
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
