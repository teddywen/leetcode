/*
https://leetcode-cn.com/problems/max-consecutive-ones/
*/
package main

import "fmt"

func main() {
	var (
		nums = []int{1, 1, 0, 1, 1, 1}
	)
	fmt.Printf("%d", findMaxConsecutiveOnes(nums))
}

// 滑动窗口
//func findMaxConsecutiveOnes(nums []int) int {
//	var (
//		res, left, right int
//	)
//	for right < len(nums) {
//		c := nums[right]
//		right++
//
//		if c == 0 || nums[left] == 0 {
//			left = right - 1
//		}
//
//		if c == 1 {
//			cnt := right - left
//			if cnt > res {
//				res = cnt
//			}
//		}
//	}
//
//	return res
//}

// 一次遍历
func findMaxConsecutiveOnes(nums []int) int {
	var (
		cnt, res int
	)
	for _, num := range nums {
		if num == 0 {
			cnt = 0
		} else {
			cnt++
		}
		if cnt > res {
			res = cnt
		}
	}
	return res
}