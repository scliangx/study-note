package main

var (
	color   []bool
	visited []bool
	ok      bool
)

// 886-可能的二分法
func possibleBipartition(n int, dislikes [][]int) bool {
	ok = true
	color = make([]bool, n+1)
	visited = make([]bool, n+1)
	graph := buildGraph(n, dislikes)
	for i := 0; i < n; i++ {
		if !visited[i] {
			bfs(graph, i)
		}
	}
	return ok
}

func bfs(graph [][]int, n int) {
	// 是否是二分图
	if !ok {
		return
	}
	// 遍历之后标记
	visited[n] = true
	for _, v := range graph[n] {
		// 当前没有遍历过
		if !visited[v] {
			// 染上一个和前一个相邻结点不一样的颜色
			color[v] = !color[n]
			bfs(graph, v)
		} else if color[v] == color[n] {
			// 两个颜色相等
			ok = false
		}
	}
}

func buildGraph(n int, dislikes [][]int) [][]int {
	graph := make([][]int, n+1)
	//  [[1,2],[1,3],[2,4]]
	for _, val := range dislikes {
		// 构造图， 相互依赖
		graph[val[0]] = append(graph[val[0]], val[1])
		graph[val[1]] = append(graph[val[1]], val[0])

	}
	return graph
}
