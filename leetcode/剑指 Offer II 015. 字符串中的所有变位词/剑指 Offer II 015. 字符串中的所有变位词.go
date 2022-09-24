package main

// 剑指 Offer II 015. 字符串中的所有变位词
func findAnagrams(s string, p string) []int {
	if len(s) == 0 {
		return []int{}
	}
	left, right, valid := 0, 0, 0
	res := []int{}
	need, window := make(map[byte]int), make(map[byte]int)
	for i := range p {
		need[p[i]]++
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
		for right-left >= len(p) {
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
