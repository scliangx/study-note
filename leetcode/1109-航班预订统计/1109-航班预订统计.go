package main

func corpFlightBookings(bookings [][]int, n int) []int {
	res := make([]int, n)
	if len(bookings) == 0 {
		return res
	}
	diff := buildDiff(res)
	for _, b := range bookings {
		i := b[0] - 1
		j := b[1] - 1
		val := b[2]
		increment(i, j, val, &diff)
	}
	// 恢复
	res[0] = diff[0]
	for i := 1; i < len(res); i++ {
		res[i] = res[i-1] + diff[i]
	}
	return res
}

// 构建查分数组
func buildDiff(nums []int) []int {
	diff := make([]int, len(nums))

	diff[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		diff[i] = nums[i] - nums[i-1]
	}
	return diff
}

// 对差分数组区间的值累加
func increment(i, j, val int, diff *[]int) {
	(*diff)[i] += val
	if j+1 < len(*diff) {
		(*diff)[j+1] -= val
	}
}
