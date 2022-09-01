package main


 
func carPooling(trips [][]int, capacity int) bool {
	res := make([]int, 1001)
	diff := difference(res)
	for _, trip := range trips {
		i := trip[1]
		j := trip[2] - 1
		val := trip[0]
		increment(i, j, val, &diff)
	}
	res[0] = diff[0]
	for i := 1; i < len(diff); i++ {
		res[i] = diff[i] + res[i-1]
	}

	for _, v := range res {
		if v > capacity {
			return false
		}
	}
	return true
}

func difference(nums []int) []int {
	diff := make([]int, len(nums))
	diff[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		diff[i] = nums[i] - nums[i-1]
	}
	return diff
}

func increment(i, j, val int, diff *[]int) {
	(*diff)[i] += val
	if j < len(*diff) {
		(*diff)[j+1] -= val
	}
	return
}
