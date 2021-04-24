/*
https://leetcode-cn.com/problems/reverse-linked-list/
*/
package main

import "fmt"

func main() {
	var (
		node1 = &ListNode{Val: 1}
		node2 = &ListNode{Val: 2}
		node3 = &ListNode{Val: 3}
		node4 = &ListNode{Val: 4}
		node5 = &ListNode{Val: 5}
	)
	node4.Next = node5
	node3.Next = node4
	node2.Next = node3
	node1.Next = node2

	last := reverseList(node1)

	for last != nil {
		fmt.Println(last.Val)
		last = last.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 递归反转法
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	last := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}

// 头删头插法
//func reverseList(head *ListNode) *ListNode {
//	var (
//		newHead *ListNode
//	)
//	for head != nil {
//		// 头删
//		node := head
//		head = head.Next
//		// 头插
//		node.Next = newHead
//		newHead = node
//	}
//	return newHead
//}
