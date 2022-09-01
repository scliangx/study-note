package main

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

// 92-反转链表(区间反转) II
var successor *ListNode

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == 1 {
		return reverseN(head, right)
	}
	head.Next = reverseBetween(head.Next, left-1, right-1)
	return head
}

func reverseN(head *ListNode, n int) *ListNode {
	if n == 1 {
		successor = head.Next
		return head
	}
	last := reverseN(head.Next, n-1)
	head.Next.Next = head
	head.Next = successor
	return last
}
