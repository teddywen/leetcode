/*
https://leetcode-cn.com/problems/binary-search/
 */
package main

import "fmt"

func main() {
	var (
		nums = []int{-1,0,3,5,9,12}
		target = 9
	)
	fmt.Printf("%#v", search(nums, target))
}

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	var (
		left = 0
		right = len(nums) - 1
	)
	for left <= right {
		mid := left + (right - left) / 2
		if nums[mid] == target {
			return mid
		} else  if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}

	return -1
}
