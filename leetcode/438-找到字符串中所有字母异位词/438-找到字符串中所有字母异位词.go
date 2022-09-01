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
	left, valid := -1, 0
	for right := 0; right < len(s); right++ {
		c := s[right]
		window[c]++
		if window[c] == need[c] {
			valid++
		}
		for right-left >= len(p) {
			// 合法添加答案即可
			if valid == len(need) {
				res = append(res, left+1)
			}
			left++
			d := s[left]
			if window[d] == need[d] {
				valid--
			}
			window[d]--
		}
	}
	return res
}
