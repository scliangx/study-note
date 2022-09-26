package main

import "math"

// 剑指 Offer II 017. 含有所有字符的最短字符串
func minWindow(s string, t string) string {
	if len(t) == 0 {
		return ""
	}
	left, right := 0, 0
	window, need := make(map[byte]int), make(map[byte]int)
	valid := 0
	start := 0
	size := math.MaxInt

	for i := range t {
		need[t[i]]++
	}
	for right < len(s) {
		c := s[right]
		right++
		if need[c] > 0 {
			window[c]++
			if need[c] == window[c] {
				valid++
			}
		}
		for valid == len(need) {
			if right-left < size {
				start = left
				size = right - left
			}
			d := s[left]
			left++
			if need[d] > 0 {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	if size == math.MaxInt {
		return ""
	}
	return s[start : start+size]
}
