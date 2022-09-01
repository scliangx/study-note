package main

import (
	"math"
)

// 76-最小覆盖子串
func minWindow(s string, t string) string {
	targetMap := make(map[string]int)
	windown := make(map[string]int)
	for i := 0; i < len(t); i++ {
		targetMap[string(t[i])]++
	}
	// 记录窗口的左边位置
	left := -1
	// 记录t中每个字符是否满足要求个数
	valid := 0
	// 记录结束位置
	end := 0
	// 开始位置和当前窗口的长度
	start, length := 0, math.MaxInt
	for right := 0; right < len(s); right++ {
		c := string(s[right])
		windown[c]++
		// 当前字符的个数窗口满足
		if windown[c] == targetMap[c] {
			valid++
		}
		// 左边位置是否需要收缩
		for valid == len(targetMap) {
			if right-left < length {
				start, end = left, right
				length = right - left
			}
			left++
			d := string(s[left])
			// 左边收缩
			if windown[d] == targetMap[d] {
				valid--
			}
			windown[d]--
		}
	}
	if length == math.MaxInt {
		return ""
	}
	return s[start+1 : end+1]
}
