package main

/**
 * Definition for singly-linked list.
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// 23-合并K个有序链表
func mergeKLists(lists []*ListNode) *ListNode {
	var pre, cur *ListNode
	for i := 0; i < len(lists); i++ {
		if i == 0 {
			pre = lists[0]
			continue
		}
		cur = lists[i]
		pre = mergeTwoLists(pre, cur)
	}
	return pre
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	dummy := &ListNode{
		Val: -1,
	}
	// p 是新的合并完成的链表的移动指针
	p, p1, p2 := dummy, list1, list2

	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			p.Next = p1
			p1 = p1.Next
		} else {
			p.Next = p2
			p2 = p2.Next
		}
		p = p.Next
	}
	if p1 == nil {
		p.Next = p2
	}
	if p2 == nil {
		p.Next = p1
	}

	return dummy.Next
}
