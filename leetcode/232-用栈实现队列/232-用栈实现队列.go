package main

// 232-用栈实现队列
type MyQueue struct {
	queue []int
}

func Constructor() MyQueue {
	return MyQueue{
		queue: []int{},
	}
}

func (this *MyQueue) Push(x int) {
	this.queue = append(this.queue, x)
}

func (this *MyQueue) Pop() int {
	if this.Empty() {
		return -1
	}
	val := this.queue[0]
	this.queue = this.queue[1:]
	return val
}

func (this *MyQueue) Peek() int {
	if this.Empty() {
		return -1
	}
	return this.queue[0]
}

func (this *MyQueue) Empty() bool {
	return len(this.queue) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
