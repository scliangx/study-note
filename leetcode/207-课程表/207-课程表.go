package main

// 207-课程表
func canFinish(numCourses int, prerequisites [][]int) bool {
	// 构建一个图
	graph := buildGraph(numCourses, prerequisites)
	count := 0
	// 记录入度
	indegree := make([]int, numCourses)
	for _, edge := range prerequisites {
		// 每一个的edge[1] 依赖于edge[0],统计edge[0]的入度
		indegree[edge[0]]++
	}
	// 如果没有依赖，则放入队列
	var queue []int
	for i := 0; i < numCourses; i++ {
		if indegree[i] == 0 {
			queue = append(queue, i)
		}
	}
	// 队列不为空的时候遍历队列
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		count++
		for _, g := range graph[cur] {
			// 遍历一个，入度减少，入度为0的时候放入队列
			indegree[g]--
			if indegree[g] == 0 {
				queue = append(queue, g)
			}
		}
	}
	// 如果最后遍历的结点个数和结点总数相等，则没有循环依赖
	return count == numCourses
}

func buildGraph(numCourses int, prerequisites [][]int) [][]int {
	graph := make([][]int, numCourses)
	for i := 0; i < numCourses; i++ {
		graph[i] = []int{}
	}
	for _, val := range prerequisites {
		graph[val[1]] = append(graph[val[1]], val[0])
	}
	return graph
}
