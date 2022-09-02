package main

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

// 83-删除排序链表中的重复元素
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	fast, slow := head, head
	for fast != nil {
		// 如果相等删除结点
		if slow.Val != fast.Val {
			slow.Next = fast
			slow = slow.Next
		}
		fast = fast.Next
	}
	// 断掉后续重复的部分
	slow.Next = nil
	return head
}
