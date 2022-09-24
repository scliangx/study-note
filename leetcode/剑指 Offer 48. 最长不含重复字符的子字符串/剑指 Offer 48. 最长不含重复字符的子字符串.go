package main

import "math"

// 剑指 Offer 48. 最长不含重复字符的子字符串
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	left, right := 0, 0
	res := math.MinInt
	window := make(map[byte]int)
	for right < len(s) {
		c := s[right]
		right++
		window[c]++
		// right 已经包含重复的了，所以left得++,得到的才是不重复的长度
		for window[c] > 1 {
			d := s[left]
			left++
			window[d]--
		}
		if res < right-left {
			res = right - left
		}
	}
	return res
}
