package main

// 986-区间列表的交集
func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	if len(firstList) == 0 || len(secondList) == 0 {
		return [][]int{}
	}
	res := [][]int{}
	i, j := 0, 0
	for i < len(firstList) && j < len(secondList) {
		a1, a2 := firstList[i][0], firstList[i][1]
		b1, b2 := secondList[j][0], secondList[j][1]
		// 有交集
		if b2 >= a1 && b1 <= a2 {
			res = append(res, []int{max(a1, b1), min(a2, b2)})
		}
		// 小的集合索引增加
		if b2 < a2 {
			j++
		} else {
			i++
		}
	}
	return res
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
