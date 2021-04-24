/*
https://leetcode-cn.com/contest/weekly-contest-235/problems/minimum-absolute-sum-difference/
*/
package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var (
		nums1 = []int{1, 10, 4, 4, 2, 7}
		nums2 = []int{9, 3, 5, 1, 7, 4}
	)
	fmt.Printf("%d", minAbsoluteSumDiff(nums1, nums2))
}

func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {
	var (
		rawASD       int
		maxReduceAD  int
		nums1Ordered = make([]int, len(nums1))
	)

	// 排序
	copy(nums1Ordered, nums1)
	sort.Sort(sort.IntSlice(nums1Ordered))

	for i := 0; i < len(nums1); i++ {
		minAD := abs5724(nums1[i] - nums2[i])
		rawAD := minAD
		rawASD += abs5724(nums1[i] - nums2[i])

		left := 0
		right := len(nums1Ordered) - 1
		for left <= right {
			mid := left + (right-left)/2
			if nums1Ordered[mid] == nums2[i] {
				minAD = 0
				break
			} else if nums1Ordered[mid] < nums2[i] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		if left >= 0 && left < len(nums1Ordered) {
			if abs5724(nums1Ordered[left]-nums2[i]) < minAD {
				minAD = abs5724(nums1Ordered[left] - nums2[i])
			}
		}
		if right >= 0 && right < len(nums1Ordered) {
			if abs5724(nums1Ordered[right]-nums2[i]) < minAD {
				minAD = abs5724(nums1Ordered[right] - nums2[i])
			}
		}

		if rawAD-minAD > maxReduceAD {
			maxReduceAD = rawAD - minAD
		}
	}

	return (rawASD - maxReduceAD) % (int(math.Pow(10, 9)) + 7)
}

func abs5724(x int) int {
	return int(math.Abs(float64(x)))
}
