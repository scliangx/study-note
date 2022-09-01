package main

// 3-无重复字符的最长子串
func lengthOfLongestSubstring(s string) int {
	window := make(map[byte]int)
	left, res := -1, 0
	for right := 0; right < len(s); right++ {
		c := s[right]
		window[c]++
		// 只需要判断收缩即可
		for window[c] > 1 {
			left++
			d := s[left]
			window[d]--
		}
		// 去当前子串和满足情况的子串的长度的最大值
		res = max(right-left, res)
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
