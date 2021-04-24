/*
https://leetcode-cn.com/problems/merge-sorted-array/
*/
package main

import "fmt"

func main() {
	var (
		nums1 = []int{1, 2, 3, 0, 0, 0}
		m     = 3
		nums2 = []int{2, 5, 6}
		n     = 3
	)
	merge(nums1, m, nums2, n)
	fmt.Printf("%+v", nums1)
}

// 逆向双指针
func merge(nums1 []int, m int, nums2 []int, n int) {
	var (
		i = m - 1
		j = n - 1
		k = m + n - 1
	)
	for i >= 0 && j >= 0 {
		if nums1[i] < nums2[j] {
			nums1[k] = nums2[j]
			j--
		} else {
			nums1[k] = nums1[i]
			i--
		}
		k--
	}
	for i >= 0 {
		nums1[k] = nums1[i]
		i--
		k--
	}
	for j >= 0 {
		nums1[k] = nums2[j]
		j--
		k--
	}
}

// 常规双指针
//func merge(nums1 []int, m int, nums2 []int, n int) {
//	var (
//		nums3   = make([]int, m+n)
//		i, j, k int
//	)
//	for i < m && j < n {
//		if nums1[i] < nums2[j] {
//			nums3[k] = nums1[i]
//			i++
//		} else {
//			nums3[k] = nums2[j]
//			j++
//		}
//		k++
//	}
//	for i < m {
//		nums3[k] = nums1[i]
//		i++
//		k++
//	}
//	for j < n {
//		nums3[k] = nums2[j]
//		j++
//		k++
//	}
//	copy(nums1, nums3)
//}
