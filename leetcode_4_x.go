/*
https://leetcode-cn.com/problems/median-of-two-sorted-arrays/
*/
package main

import "fmt"

func main() {
	var (
		nums1 = []int{1, 3}
		nums2 = []int{2}
	)
	fmt.Printf("%#v", findMedianSortedArrays(nums1, nums2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var (
		left = (len(nums1) + len(nums2)) / 2
		i, j int
	)
	for i != len(nums1) && j != len(nums2) && left > 0 {
		if nums1[i+left/2] < nums2[j+left/2] {
			//j +=
		}
	}


}
