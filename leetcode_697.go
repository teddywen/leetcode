/*
https://leetcode-cn.com/problems/degree-of-an-array/
 */
package main

import "fmt"

func main() {
	var (
		nums = []int{1,2,2,3,1,4,2}
	)
	fmt.Printf("%d", findShortestSubArray(nums))
}

// 滑动窗口
func findShortestSubArray(nums []int) int {
	var (
		degree = arrayDegree(nums)
		cnt = make(map[int]int)
		left, right, res int
	)

	for right < len(nums) {
		c := nums[right]
		right++

		cnt[c]++

		var shrink bool
		if cnt[c] == degree {
			shrink = true
		}

		for cnt[c] == degree {
			c := nums[left]
			left++

			cnt[c]--
		}

		if shrink {
			if res > 0 {
				res = min697(res, right - left + 1)
			} else {
				res = right - left + 1
			}
		}
	}

	return res
}

func arrayDegree(nums []int) int {
	var (
		cnt = make(map[int]int)
		degree int
	)

	for _, v := range nums {
		cnt[v]++
		degree = max697(degree, cnt[v])
	}
	return degree
}

func max697(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func min697(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}