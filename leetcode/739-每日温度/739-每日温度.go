package main

// 739-每日温度
func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	stack := []int{}
	for i := len(temperatures) - 1; i >= 0; i-- {
		num := temperatures[i]
		for len(stack) > 0 && num >= temperatures[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		// 计算索引的差值
		if len(stack) > 0 {
			res[i] = stack[len(stack)-1] - i
		} else {
			res[i] = 0
		}
		// 栈中记录索引的值
		stack = append(stack, i)
	}

	return res
}
