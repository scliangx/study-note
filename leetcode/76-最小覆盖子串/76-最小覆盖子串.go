package main

import (
	"math"
)

// 76-最小覆盖子串
func minWindow(s string, t string) string {
	need := make(map[string]int)
	window := make(map[string]int)
	for i := 0; i < len(t); i++ {
		need[string(t[i])]++
	}
	// 记录窗口的左右边位置
	left, right := 0, 0
	// 记录t中每个字符是否满足要求个数
	valid := 0
	// 开始位置和当前窗口的长度
	start, length := 0, math.MaxInt
	for right < len(s) {
		c := string(s[right])
		right++
		if need[c] > 0 {
			window[c]++
			// 当前字符的个数窗口满足
			if window[c] == need[c] {
				valid++
			}
		}

		// 左边位置是否需要收缩
		for valid == len(need) {
			if right-left < length {
				start = left
				length = right - left
			}
			d := string(s[left])
			left++
			if need[d] > 0 {
				// 左边收缩
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	if length == math.MaxInt {
		return ""
	}
	return s[start : start+length]
}
