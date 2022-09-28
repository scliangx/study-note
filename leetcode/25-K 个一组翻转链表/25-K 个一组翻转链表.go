package main


/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

// 25-K 个一组翻转链表
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	p1, p2 := head, head
	// 不⾜ k 个，不需要反转，base case
	for i := 0; i < k; i++ {
		if p2 == nil {
			return head
		}
		p2 = p2.Next
	}
	newHead := reverseLinkList(p1, p2)
	p1.Next = reverseKGroup(p2,k)
	return newHead
}

func reverseLinkList(p1, p2 *ListNode) *ListNode {
	if p1 == nil {
		return nil
	}
	var pre *ListNode
	cur := p1
	for cur != p2 {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
