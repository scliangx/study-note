package main

// 剑指 Offer II 016. 不含重复字符的最长子字符串
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	left, right := 0, 0
	window := make(map[byte]int)
	res := 0
	for right < len(s) {
		c := s[right]
		right++
		window[c]++
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
