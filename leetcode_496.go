/*
https://leetcode-cn.com/problems/next-greater-element-i/
*/
package main

import "fmt"

func main() {
	var (
		nums1 = []int{4, 1, 2}
		nums2 = []int{1, 3, 4, 2}
	)
	fmt.Printf("%#v", nextGreaterElement(nums1, nums2))
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	var (
		nums1Dict       = make(map[int]int, len(nums1))
		monotonousStack = make([]int, 0)
	)
	for i, v := range nums1 {
		nums1Dict[v] = i
	}
	for i := len(nums2) - 1; i >= 0; i-- {
		v := nums2[i]
		for len(monotonousStack) > 0 && v >= monotonousStack[len(monotonousStack)-1] {
			monotonousStack = monotonousStack[:len(monotonousStack)-1]
		}

		if pos, ok := nums1Dict[v]; ok {
			if len(monotonousStack) == 0 {
				nums1[pos] = -1
			} else {
				nums1[pos] = monotonousStack[len(monotonousStack)-1]
			}
		}

		monotonousStack = append(monotonousStack, v)
	}
	return nums1
}
