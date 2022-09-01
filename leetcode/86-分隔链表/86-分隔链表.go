package main

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

// 86-分隔链表 
func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}
	dummy1,dummy2 := &ListNode{},&ListNode{}
	p1,p2,p := dummy1,dummy2,head
	for p != nil{
		if p.Val < x {
			p1.Next = p
			p1 = p1.Next
		}else{
			p2.Next = p
			p2 = p2.Next
		}
		temp := p.Next
		p.Next = nil
		p = temp
	}
	p1.Next = dummy2.Next
	return dummy1.Next
}
