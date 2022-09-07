package main

// 39-组合总和
func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	track := []int{}
	if len(candidates) == 0 {
		return res
	}
	backtrack(candidates, 0, 0, target, track, &res)
	return res
}

func backtrack(nums []int, start, sumVal, target int, track []int, res *[][]int) {
	if sumVal == target {
		tmp := make([]int, len(track))
		copy(tmp, track)
		*res = append(*res, tmp)
		return
	}
	if sumVal > target {
		return
	}
	for i := start; i < len(nums); i++ {
		sumVal += nums[i]
		track = append(track, nums[i])
		backtrack(nums, i, sumVal, target, track, res)
		track = track[:len(track)-1]
		sumVal -= nums[i]
	}
}
