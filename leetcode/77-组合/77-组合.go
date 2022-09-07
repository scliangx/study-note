package main

// 77-组合
func combine(n int, k int) [][]int {
	res := [][]int{}
	track := []int{}
	if n == 0 {
		return res
	}
	backtrack(n, k, 0, &res, &track)
	return res
}

func backtrack(n int, k int, start int, res *[][]int, track *[]int) {
	if len(*track) == k {
		tmp := make([]int, len(*track))
		copy(tmp, *track)
		*res = append(*res, tmp)
		return
	}
	for i := start; i < n; i++ {
		*track = append(*track, i)
		backtrack(n, k, i+1, res, track)
		*track = (*track)[:len(*track)-1]
	}
}
