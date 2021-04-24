/*
https://leetcode-cn.com/problems/longest-increasing-subsequence/
*/
package main

import "fmt"

func main() {
	var (
		nums = []int{10,9,2,5,3,7,101,18}
	)
	fmt.Printf("%d", lengthOfLIS(nums))
}

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return 1
	}

	res := 1
	// dp(i): 以第i个元素结尾的最长子序列
	dp := make([]int, len(nums))
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		j := -1 // j为以i为下标的元素的上一个lis元素的下标, dp(i) = dp(j) + 1
		for k := 0; k < i; k++ {
			if nums[k] >= nums[i] {
				continue
			}
			if j == -1 {
				j = k
			} else {
				if dp[k] > dp[j] {
					j = k
				}
			}
		}
		if j == -1 {
			dp[i] = 1
		} else {
			dp[i] = dp[j] + 1
		}
		if res < dp[i] {
			res = dp[i]
		}
	}
	return res
}
