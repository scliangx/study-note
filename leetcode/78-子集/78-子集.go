package main

// 78-子集
func subsets(nums []int) [][]int {
	res := [][]int{}
	track := []int{}
	if len(nums) == 0 {
		return res
	}
	backtrack(nums, &res, &track, 0)
	return res
}

func backtrack(nums []int, res *[][]int, track *[]int, start int) {
	tmp := make([]int, len(*track))
	copy(tmp, *track)
	*res = append(*res, tmp)
	for i := start; i < len(nums); i++ {
		*track = append(*track, nums[i])
		backtrack(nums, res, track, i+1)
		*track = (*track)[:len(*track)-1]
	}
}
