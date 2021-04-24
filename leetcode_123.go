/*
https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii/
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		nums = []int{3, 3, 5, 0, 0, 3, 1, 4}
	)
	fmt.Printf("%d", maxProfit(nums))
}

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	var (
		// dp[i][j][k] i:第i+1天 j: 0表示手里没股票,1表示手里有股票 k:还能完成k+1笔交易
		dp  = make([][][]int, len(prices))
		res int
	)
	dp[0] = make([][]int, 2)
	dp[0][1] = []int{-prices[0], -prices[0], -prices[0]}
	dp[0][0] = []int{0, 0, 0}

	for i := 1; i < len(prices); i++ {
		dp[i] = make([][]int, 2)
		dp[i][1] = make([]int, 3)
		dp[i][0] = make([]int, 3)
		for k := 2; k >= 0; k-- {
			if k == 2 {
				dp[i][0][k] = dp[i-1][0][k]
			} else {
				dp[i][0][k] = int(math.Max(
					float64(dp[i-1][0][k]),
					float64(dp[i-1][1][k+1]+prices[i]),
				))
			}
			dp[i][1][k] = int(math.Max(
				float64(dp[i-1][1][k]),
				float64(dp[i-1][0][k]-prices[i]),
			))
		}
	}

	for i := 0; i <= 2; i++ {
		if res < dp[len(prices)-1][0][i] {
			res = dp[len(prices)-1][0][i]
		}
	}
	return res
}
