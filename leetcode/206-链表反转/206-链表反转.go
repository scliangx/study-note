package main

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

// 206-链表反转
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var prev *ListNode
	cur := head
	for cur != nil{
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}
