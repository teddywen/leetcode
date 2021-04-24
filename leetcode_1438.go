/*
https://leetcode-cn.com/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit/
*/
package main

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {
	var (
		nums  = []int{8, 2, 4, 7}
		limit = 4
	)
	fmt.Printf("%#v", longestSubarray(nums, limit))
}

func longestSubarray(nums []int, limit int) int {
	var (
		left, right, min, max int
	)
	for right < len(nums) {
		c := nums[right]
		right++

		min = min1438(min, c)
		max = max1438(max, c)

		for max-min > limit {
			c := nums[left]
			left++
			max = min1438(max, c)
		}
	}
}

func min1438(x, y int) int {
	// 因为 1 <= nums[i] <= 10^9
	if x == 0 {
		return y
	}
	if x < y {
		return x
	} else {
		return y
	}
}

func max1438(x, y int) int {
	// 因为 1 <= nums[i] <= 10^9
	if x == 0 {
		return y
	}
	if x > y {
		return x
	} else {
		return y
	}
}

// An IntHeap1438 is a min-heap of ints.
type IntHeap1438 []int

func (h IntHeap1438) Len() int           { return len(h) }
func (h IntHeap1438) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap1438) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap1438) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap1438) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
