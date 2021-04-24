/*
https://leetcode-cn.com/contest/weekly-contest-231/problems/make-the-xor-of-all-segments-equal-to-zero/
*/
package main

import (
	"fmt"
)

func main() {
	var (
		nums = []int{26, 19, 19, 28, 13, 14, 6, 25, 28, 19, 0, 15, 25, 11}
		k    = 3
	)
	fmt.Printf("%d", minChanges(nums, k))
}

// 贪心试试
func minChanges(nums []int, k int) int {
	return _minChanges(nums, k)
}

func _minChanges(nums []int, k int) int {
	var (
		result, xor int
		numsChange = make([]int, len(nums))
	)
	copy(numsChange, nums)
	for i := 0; i < k-1; i++ {
		xor ^= numsChange[i]
	}
	for i := k - 1; i < len(numsChange); i++ {
		if numsChange[i] != xor {
			result++

			numsChange[i] = xor
		}
		//xor = numsChange[i]
		for j := 0; j < k-2; j++ {
			xor ^= numsChange[i-j-1]
		}
	}
	return result
}
