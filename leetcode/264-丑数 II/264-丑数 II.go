package main

// 264-丑数 II
func nthUglyNumber(n int) int {
	if n == 0 {
		return 0
	}
	// 相当于2,3,5的倍数组成的三条链表，合并去重即可
	p2, p3, p5 := 1, 1, 1
	headP2, headP3, headP5 := 1, 1, 1
	unly := make([]int, n+1)
	pre := 1
	for pre <= n {
		minVal := min(min(headP2, headP3), headP5)
		unly[pre] = minVal
		pre++
		if minVal == headP2 {
			headP2 = 2 * unly[p2]
			p2++
		}
		if minVal == headP3 {
			headP3 = 3 * unly[p3]
			p3++
		}
		if minVal == headP5 {
			headP5 = 5 * unly[p5]
			p5++
		}
	}
	return unly[n]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
