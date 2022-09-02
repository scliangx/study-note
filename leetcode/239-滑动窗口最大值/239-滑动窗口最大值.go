package main

type MonotonicQueue struct {
	queue []int
}

// 构建单调队列
func makeMonotonicQueue() *MonotonicQueue {
	return &MonotonicQueue{
		queue: []int{},
	}
}

// push
func (queue *MonotonicQueue) push(val int) {
	// 如果队列不为空，需要添加的元素比队尾的元素小，则直接删除队尾的元素
	for len(queue.queue) != 0 && queue.queue[len(queue.queue)-1] < val {
		queue.queue = queue.queue[:len(queue.queue)-1]
	}
	// 遇到比自己大的添加到队尾
	queue.queue = append(queue.queue, val)
}

// 对头的就是最大的元素
func (queue *MonotonicQueue) max() int {
	return queue.queue[0]
}

// pop
func (queue *MonotonicQueue) pop(val int) {
	if queue.queue[0] == val {
		queue.queue = queue.queue[1:]
	}
}

// 239-滑动窗口最大值
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	queue := makeMonotonicQueue()
	res := []int{}
	for i := 0; i < len(nums); i++ {
		if i < k-1 {
			//先把窗⼝的前 k - 1 填满
			queue.push(nums[i])
		} else {
			// 窗⼝开始向前滑动
			// 移⼊新元素
			queue.push(nums[i])
			// 将当前窗⼝中的最⼤元素记⼊结果
			res = append(res, queue.max())
			// 当前i->k+1部分找到最大值后，移出对头的元素
			queue.pop(nums[i-k+1])
		}
	}
	return res
}
