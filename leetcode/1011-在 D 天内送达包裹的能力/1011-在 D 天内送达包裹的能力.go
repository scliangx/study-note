package main

// 1011-在 D 天内送达包裹的能力
func shipWithinDays(weights []int, days int) int {
	if len(weights) == 0 {
		return 0
	}
	// 对于左边界而言，由于我们不能「拆分」一个包裹，因此船的运载能力不能小于所有包裹中最重的那个的重量，即左边界为数组 weights 中元素的最大值
	// 对于右边界而言，船的运载能力也不会大于所有包裹的重量之和，即右边界为数组 \textit{weights}weights 中元素的和。
	left, right := 0, 1
	for _, v := range weights {
		left = max(left, v)
		right += v
	}
	for left < right {
		mid := left + (right-left)/2
		if f(weights, mid) <= days {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

// 定义：当运载能⼒为 x 时，需要 f(x) 天运完所有货物
// f(x) 随着 x 的增加单调递减
func f(weights []int, x int) int {
	days := 0
	for i := 0; i < len(weights); {
		// 尽可能多的装货物
		cap := x
		for i < len(weights) {
			// 当前运载能力小于当前货物的重量，下一天
			if cap < weights[i] {
				break
			} else {
				// 装运，cap - weights[i]
				cap -= weights[i]
				i++
			}
		}
		days++
	}
	return days
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
