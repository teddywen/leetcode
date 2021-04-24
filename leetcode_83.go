/*
https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/
*/
package main

import (
	"fmt"
	"math"
)

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

func deleteDuplicates(head *ListNode) *ListNode {
	var (
		newTail = head
		curHead = head
	)
	for curHead != nil {

		node := curHead
		curHead = curHead.Next
		node.Next = nil

		if newTail.Val != node.Val {
			newTail.Next = node
			newTail = node
		}
	}
	return head
}
