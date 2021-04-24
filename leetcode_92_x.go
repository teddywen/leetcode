/*
https://leetcode-cn.com/problems/reverse-linked-list-ii/
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
		left  = 2
		right = 4
	)
	node4.Next = node5
	node3.Next = node4
	node2.Next = node3
	node1.Next = node2

	last := reverseBetween(node1, left, right)

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
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	return nil
}

// 头删头插法
//func reverseBetween(head *ListNode, left int, right int) *ListNode {
//
//}
