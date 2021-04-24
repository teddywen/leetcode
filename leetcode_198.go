/*
https://leetcode-cn.com/problems/house-robber/
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		nums = []int{1, 2, 3, 1}
	)
	fmt.Printf("%d", rob(nums))
}

//func rob(nums []int) int {
//	var (
//		// 定义 dp[i][j] = x
//		// i: 打劫到第i家人
//		// j: 当天是否打劫 0表示当天未打劫 1表示当天打劫了
//		// x: 最大收益
//		dp = make([][2]int, len(nums)+1)
//	)
//	for i := 1; i <= len(nums); i++ {
//		// 当天不打劫
//		dp[i][0] = dp[i-1][0]
//		if dp[i][0] < dp[i-1][1] {
//			dp[i][0] = dp[i-1][1]
//		}
//		// 当天打劫
//		dp[i][1] = dp[i-1][0] + nums[i-1]
//	}
//	return int(math.Max(
//		float64(dp[len(nums)][0]),
//		float64(dp[len(nums)][1]),
//	))
//}

// 路径压缩版
func rob(nums []int) int {
	var (
		// 定义 dp[i][j] = x
		// i: 打劫到第i家人
		// j: 当天是否打劫 0表示当天未打劫 1表示当天打劫了
		// x: 最大收益
		//dp = make([][2]int, len(nums)+1)
		last, cur = [2]int{}, [2]int{}
	)
	for i := 1; i <= len(nums); i++ {
		// 当天不打劫
		cur[0] = last[0]
		if cur[0] < last[1] {
			cur[0] = last[1]
		}
		// 当天打劫
		cur[1] = last[0] + nums[i-1]
		last = cur
	}
	return int(math.Max(
		float64(cur[0]),
		float64(cur[1]),
	))
}
