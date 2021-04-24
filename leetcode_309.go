/*
https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		nums = []int{1, 2, 3, 0, 2}
	)
	fmt.Printf("%d", maxProfit(nums))
}

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	var (
		dp  = make([][][]int, len(prices))
		res int
		cd  = 1
	)
	dp[0] = make([][]int, 2)
	dp[0][0] = make([]int, cd+1)
	dp[0][1] = make([]int, cd+1)
	for j := 0; j < cd+1; j++ {
		dp[0][1][j] = -prices[0]
	}

	// dp[i][0][j] 手里没股票的最大利润 i:第i+1天 j:当前冷冻期j天
	// dp[i][1][j] 手里有股票的最大利润 i:第i+1天 j:当前冷冻期j天
	for i := 1; i < len(prices); i++ {
		dp[i] = make([][]int, 2)
		dp[i][1] = make([]int, cd+1)
		dp[i][0] = make([]int, cd+1)
		for j := cd; j >= 0; j-- {

			if j < cd {
				// 前一天有冷冻期
				dp[i][0][j] = dp[i-1][0][j+1]
				// 前一天没冷冻期
				if j == 0 {
					dp[i][0][j] = int(math.Max(
						float64(dp[i-1][0][j]),
						float64(dp[i][0][j]),
					))
					dp[i][1][j] = int(math.Max(
						float64(dp[i-1][0][j]-prices[i]),
						float64(dp[i-1][1][j]),
					))
				}
			} else {
				// 当天卖出，冷冻期变为k
				dp[i][0][j] = dp[i-1][1][0] + prices[i]
			}
		}
	}

	for j := 0; j < cd+1; j++ {
		if res < dp[len(prices)-1][0][j] {
			res = dp[len(prices)-1][0][j]
		}
	}
	return res
}
