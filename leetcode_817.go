/*
https://leetcode-cn.com/problems/linked-list-components/
 */
package main

func main() {
	//var (
	//	nums = []int{2,7,11,15}
	//	target = 9
	//)
	//fmt.Printf("%d", numComponents(nums, target))
}

func numComponents(head *ListNode, G []int) int {
	var (
		res, max int
		visited = make(map[int]struct{}, len(G))
	)
	for _, g := range G {
		visited[g] = struct{}{}
	}

	for head != nil {
		if _, ok := visited[head.Val]; ok {
			if max == 0 {
				res++
			}
			max++
		} else {
			max = 0
		}
		head = head.Next
	}

	return res
}

// Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}