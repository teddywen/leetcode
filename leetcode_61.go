/*
https://leetcode-cn.com/problems/rotate-list/
*/
package main

import "fmt"

func main() {
	var (
		nums   = []int{2, 7, 11, 15}
		target = 9
	)
	fmt.Printf("%#v", twoSum(nums, target))
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	var (
		slow = head
		fast = head
		pre = head
		count = head
		n int
	)
	for count != nil {
		count = count.Next
		n++
	}
	k = k % n - 1
	for i := 1; i < k; i++ {
		if fast.Next == nil {
			fast = head
		} else {
			fast = fast.Next
		}
	}
	for fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next
	}
	if slow == head {
		return head
	}
	pre.Next = nil
	fast.Next = head
	head = slow
	return slow
}
