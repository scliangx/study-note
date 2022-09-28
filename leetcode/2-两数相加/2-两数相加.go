package main

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

// 2-两数相加
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	dummy := &ListNode{Val: -999}
	newHead := dummy
	carry := 0
	// 两个同时不为0,并且进位不为0
	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		// 进位累加到下个结点求和
		carry = sum / 10
		sum %= 10
		newHead.Next = &ListNode{Val: sum}
		newHead = newHead.Next
	}
	return dummy.Next
}
