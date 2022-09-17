package main

// 392-判断子序列
func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	curS, curT := 0, 0
	for curS < len(s) && curT < len(t) {
		if s[curS] == t[curT] {
			curS++
		}
		curT++
	}

	return curS == len(s)
}
