/*
https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii/
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
		newHead, newTail *ListNode
		curHead = head
		repeat = math.MinInt32
	)
	for curHead != nil {

		node := curHead
		curHead = curHead.Next

		if node.Next !=nil && node.Val == node.Next.Val {
			repeat = node.Val
		} else if node.Val != repeat {
			if newTail != nil {
				newTail.Next = node
			}
			node.Next = nil
			newTail = node
			if newHead == nil {
				newHead = node
			}
		}
	}
	return newHead
}
