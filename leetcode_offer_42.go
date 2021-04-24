/*
https://leetcode-cn.com/problems/lian-xu-zi-shu-zu-de-zui-da-he-lcof/
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		nums = []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	)
	fmt.Printf("%d", maxSubArray(nums))
}

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var dp = make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1] <= 0 {
			dp[i] = nums[i]
		} else {
			dp[i] = dp[i-1] + nums[i]
		}
	}

	var res = math.MinInt32
	for _, v := range dp {
		if v > res {
			res = v
		}
	}
	return res
}
