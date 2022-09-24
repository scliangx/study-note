package main

// 438-找到字符串中所有字母异位词
func findAnagrams(s string, p string) []int {
	need := make(map[byte]int)
	window := make(map[byte]int)
	// 记录下标索引，返回答案
	res := []int{}
	for i := range p {
		need[p[i]]++
	}
	left, right, valid := 0, 0, 0
	for right < len(s) {
		c := s[right]
		right++
		if need[c] > 0 {
			window[c]++
			// 有一个字符满足要求，则valid增加1
			if window[c] == need[c] {
				valid++
			}
		}
		// left滑动
		for right-left >= len(p) {
			// 所有都满足要求之后进行添加left到res
			if valid == len(need) {
				res = append(res, left)
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
	return res
}
