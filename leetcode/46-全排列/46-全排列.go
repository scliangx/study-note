package main

// 46-全排列 回溯算法第一题
func permute(nums []int) [][]int {
	res := [][]int{}
	visited := make([]bool, len(nums))
	if nums == nil || len(nums) == 0 {
		return res
	}
	track := []int{}
	backtrack(nums, &res, track, visited)
	return res
}

func backtrack(nums []int, res *[][]int, track []int, visited []bool) {
	if len(track) == len(nums) {
		// 需要进行copy每一次的结果
		tmp := make([]int, len(track))
		copy(tmp, track)
		*res = append(*res, tmp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if visited[i] {
			continue
		}
		track = append(track, nums[i])
		visited[i] = true
		backtrack(nums, res, track, visited)
		track = track[:len(track)-1]
		visited[i] = false
	}
}
