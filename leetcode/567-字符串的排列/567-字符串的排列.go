package main

// 567-字符串的排列
func checkInclusion(s1 string, s2 string) bool {
	s1Map := make(map[byte]int, 0)
	window := make(map[byte]int, 0)
	for i := range s1 {
		s1Map[s1[i]]++
	}
	left, valid := -1, 0
	for right := 0; right < len(s2); right++ {
		c := s2[right]
		window[c]++
		if window[c] == s1Map[c] {
			valid++
		}
		// 滑动窗口太小
		if right-left < len(s1) {
			continue
		}
		// 判断是否符合条件
		if valid == len(s1Map) {
			return true
		}
		left++
		d := s2[left]

		if window[d] == s1Map[d] {
			valid--
		}
		window[d]--
	}
	return false
}
