package main

// 752-打开转盘锁
func openLock(deadends []string, target string) int {
	count := 0
	// 记录死锁黑名单
	deadendsMap := make(map[string]bool)
	// 记录已经访问过的
	visited := make(map[string]bool)
	for _, v := range deadends {
		deadendsMap[v] = true
	}
	var queue []string
	visited["0000"] = true
	queue = append(queue, "0000")
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			lock := queue[0]
			queue = queue[1:]
			if deadendsMap[lock] {
				continue
			}
			if lock == target {
				return count
			}
			for j := 0; j < 4; j++ {
				// 对四个位置每一个位置进行上下转动
				up := plusOne(lock, j)
				if !visited[up] {
					queue = append(queue, up)
					visited[up] = true
				}
				down := minusOne(lock, j)
				if !visited[down] {
					queue = append(queue, down)
					visited[down] = true
				}
			}

		}
		count++
	}
	return -1
}

// 将 s[j] 向上拨动⼀次
func plusOne(s string, i int) string {
	c := []int32{}
	for _, v := range s {
		c = append(c, v)
	}
	if c[i] == '9' {
		c[i] = '0'
	} else {
		c[i] += 1
	}
	res := ""
	for _, v := range c {
		res += string(v)
	}
	return res

}

// 将 s[j] 向下拨动⼀次
func minusOne(s string, i int) string {
	c := []int32{}
	for _, v := range s {
		c = append(c, v)
	}
	if c[i] == '0' {
		c[i] = '9'
	} else {
		c[i] -= 1
	}
	res := ""
	for _, v := range c {
		res += string(v)
	}
	return res
}
