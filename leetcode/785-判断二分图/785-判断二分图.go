package main

// 
func isBipartite(graph [][]int) bool {
	visited := make([]int, len(graph))

	for i := 0; i < len(graph); i++ {
		if visited[i] != 0 {
			continue
		}
		var queue []int
		queue = append(queue, i)
		// 等于0已经上完一种颜色
		visited[i] = 1
		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]
			for _, g := range graph[cur] {
				// 如果没有访问过
				if visited[g] == 0 {
					// 添加一个不一样的颜色
					visited[g] = -visited[cur]
					queue = append(queue, g)
				} else if visited[g] == visited[cur] {
					// 没有访问过，表示已经染过颜色，颜色和相邻的结点颜色相同，不能完全着色，返回false
					return false
				}
			}
		}
	}
	return true
}
