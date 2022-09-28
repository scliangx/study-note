package main


/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}


// 剑指 Offer 18. 删除链表的节点
func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{Val: -999}
	p := dummy
	cur := head
	for cur != nil {
		if cur.Val != val {
			p.Next = &ListNode{Val: cur.Val}
			p = p.Next
		}
		cur = cur.Next
	}
	return dummy.Next
}

func deleteNode1(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{Val: -999}
	p := dummy
	cur := head
	for cur != nil {
		if cur.Val == val {
			p.Next = cur
			p = p.Next
		}
		// 这儿必须让cur的next指向nil，不然上述操作相当于将cur后边的整个链表添加到p
		tmp := cur.Next
		cur.Next = nil
		cur = tmp
	}
	return dummy.Next
}
