/*
https://leetcode-cn.com/problems/kth-largest-element-in-an-array/?utm_source=LCUS&utm_medium=ip_redirect&utm_campaign=transfer2china
*/
package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	var (
		nums = []int{3, 2, 1, 5, 6, 4}
		k    = 2
	)
	fmt.Printf("%#v", findKthLargest(nums, k))
}

// heap
// O(NLogN)
func findKthLargest(nums []int, k int) int {
	var (
		hp = make(HeapSlice, 0)
	)
	for _, n := range nums {
		heap.Push(&hp, n)
		if len(hp) > k {
			heap.Pop(&hp)
		}
	}
	return heap.Pop(&hp).(int)
}

type HeapSlice sort.IntSlice

func (p *HeapSlice) Len() int           { return len(*p) }
func (p *HeapSlice) Less(i, j int) bool { return (*p)[i] < (*p)[j] }
func (p *HeapSlice) Swap(i, j int)      { (*p)[i], (*p)[j] = (*p)[j], (*p)[i] }
func (p *HeapSlice) Push(x interface{}) { *p = append(*p, x.(int)) }
func (p *HeapSlice) Pop() interface{} {
	x := (*p)[len(*p)-1]
	*p = (*p)[:len(*p)-1]
	return x
}

//// quick select
//// O(n)
//func findKthLargest(nums []int, k int) int {
//	shuffle(nums)
//	var (
//		lo, hi = 0, len(nums) - 1
//	)
//	k = len(nums) - k
//	for {
//		p := partition(nums, lo, hi)
//		if p < k {
//			lo = p + 1
//		} else if p > k {
//			hi = p - 1
//		} else {
//			return nums[p]
//		}
//	}
//}
//
//func partition(nums []int, lo, hi int) int {
//	if lo >= hi {
//		return lo
//	}
//	var (
//		i, j = lo, hi + 1
//	)
//	for {
//		for i++; nums[i] < nums[lo]; i++ {
//			if i == hi {
//				break
//			}
//		}
//		for j--; nums[j] > nums[lo]; j-- {
//			if j == lo {
//				break
//			}
//		}
//		if i >= j {
//			break
//		}
//		nums[i], nums[j] = nums[j], nums[i]
//	}
//	nums[j], nums[lo] = nums[lo], nums[j]
//	return j
//}
//
//func shuffle(nums []int) {
//	for i := range nums {
//		j := rand.Intn(len(nums)-i) + i
//		nums[i], nums[j] = nums[j], nums[i]
//	}
//}
