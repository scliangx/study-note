package main

type LRUCache struct {
	Cap  int
	Map  map[int]*Node
	Head *Node
	Last *Node
}

type Node struct {
	Val  int
	Key  int
	Pre  *Node
	Next *Node
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		Cap:  capacity,
		Map:  make(map[int]*Node, capacity),
		Head: &Node{},
		Last: &Node{},
	}
	lru.Head.Next = lru.Last
	lru.Last.Pre = lru.Head
	return lru

}

// 146-LRU 缓存(LRU (最近最少使用) 缓存)
/*
1. 使用一个双向链表，其中使用map存储元素
2. 每一次添加元素在头部添加
3. 使用get之后，将元素结点移动到头部位置
4. 如果删除直接删除位置
*/

func (this *LRUCache) Get(key int) int {
	node, ok := this.Map[key]
	if !ok {
		return -1
	}
	this.remove(node)
	this.addHeader(node)
	return node.Val
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.Map[key]
	if ok {
		this.remove(node)
	} else {
		if len(this.Map) == this.Cap {
			delete(this.Map, this.Last.Pre.Key)
			this.remove(this.Last.Pre)
		}
		node = &Node{Val: value, Key: key}
		this.Map[node.Key] = node
	}
	node.Val = value
	this.addHeader(node)
}

func (this *LRUCache) addHeader(node *Node) {
	this.Head.Next.Pre = node
	node.Next = this.Head.Next
	this.Head.Next = node
	node.Pre = this.Head
}

func (this *LRUCache) remove(node *Node) {
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
