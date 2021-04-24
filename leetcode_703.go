/*
https://leetcode-cn.com/problems/kth-largest-element-in-a-stream/
*/
package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	var (
		nums = []int{4, 5, 8, 2}
		k    = 3
	)
	kl := Constructor(k, nums)
	fmt.Printf("%d\n", kl.Add(3))
	fmt.Printf("%d\n", kl.Add(5))
	fmt.Printf("%d\n", kl.Add(10))
	fmt.Printf("%d\n", kl.Add(9))
	fmt.Printf("%d\n", kl.Add(4))
}

type KthLargest struct {
	sort.IntSlice
	k int
}

func Constructor(k int, nums []int) KthLargest {
	kl := KthLargest{k: k}
	for _, val := range nums {
		kl.Add(val)
	}
	return kl
}

func (kl *KthLargest) Push(v interface{}) {
	kl.IntSlice = append(kl.IntSlice, v.(int))
}

func (kl *KthLargest) Pop() interface{} {
	a := kl.IntSlice
	v := a[len(a)-1]
	kl.IntSlice = a[:len(a)-1]
	return v
}

// 用小顶堆来实现
func (kl *KthLargest) Add(val int) int {
	heap.Push(kl, val) // append到最后一个元素，然后将其往上堆化
	if kl.Len() > kl.k {
		heap.Pop(kl) // 第0个和最后一个元素交换，然后将最后一个元素去掉。将堆顶元素往下堆化
	}
	return kl.IntSlice[0] // 堆超过k时，将最小的顶掉后，堆顶还是第k大元素
}
