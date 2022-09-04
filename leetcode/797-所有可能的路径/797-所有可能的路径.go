package main

// 797-所有可能的路径
func allPathsSourceTarget(graph [][]int) [][]int {
	path := []int{}
	res := [][]int{}
	traverse(graph, 0, path, &res)
	return res
}

func traverse(graph [][]int, node int, path []int, res *[][]int) {
	path = append(path, node)
	n := len(graph) - 1
	// 如果到了终点
	if n == node {
		tmp := make([]int, len(path))
		copy(tmp, path)
		(*res) = append((*res), tmp)
		return
	}
	// 遍历每一个邻接点
	for _, v := range graph[node] {
		traverse(graph, v, path, res)
	}
	path = path[:len(path)-1]
}
