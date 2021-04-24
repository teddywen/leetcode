/*
https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		nums = []int{7, 1, 5, 3, 6, 4}
	)
	fmt.Printf("%d", maxProfit(nums))
}

// 常规解法
//func maxProfit(prices []int) int {
//	if len(prices) <= 1 {
//		return 0
//	}
//
//	var (
//		profit int
//	)
//
//	for i:=1; i<len(prices); i++ {
//		if prices[i] > prices[i-1] {
//			// 有得赚
//			profit += prices[i] - prices[i-1]
//		}
//	}
//	return profit
//}

// 动规解法
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	var (
		dp = make([][]int, len(prices))
	)
	dp[0] = make([]int, 2)
	dp[0][1] = -prices[0]
	dp[0][0] = 0

	// dp[i][0] 手里没股票的最大利润
	// dp[i][1] 手里有股票的最大利润
	for i := 1; i < len(prices); i++ {
		dp[i] = make([]int, 2) // 初始化二维状态数组
		dp[i][0] = int(math.Max(
			float64(dp[i-1][0]),            // 前一天手里没股票
			float64(dp[i-1][1]+prices[i])), // 前一天手里有股票
		)
		dp[i][1] = int(math.Max(
			float64(dp[i-1][1]),
			float64(dp[i-1][0]-prices[i]),
		))
	}
	return dp[len(prices)-1][0]
}
