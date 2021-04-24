/*
https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iv/
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		nums = []int{2, 4, 1}
		k    = 2
	)
	fmt.Printf("%d", maxProfit(k, nums))
}

func maxProfit(k int, prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	var (
		// dp[j][n][j] j:第i+1天 n: 0表示手里没股票,1表示手里有股票 j:还能完成j+1笔交易
		dp  = make([][][]int, len(prices))
		res int
	)
	dp[0] = make([][]int, 2)
	dp[0][0] = make([]int, k+1)
	dp[0][1] = make([]int, k+1)
	for j := 0; j < k+1; j++ {
		dp[0][1][j] = -prices[0]
	}

	for i := 1; i < len(prices); i++ {
		dp[i] = make([][]int, 2)
		dp[i][1] = make([]int, k+1)
		dp[i][0] = make([]int, k+1)
		for j := k; j >= 0; j-- {
			if j == k {
				dp[i][0][j] = dp[i-1][0][j]
			} else {
				dp[i][0][j] = int(math.Max(
					float64(dp[i-1][0][j]),
					float64(dp[i-1][1][j+1]+prices[i]),
				))
			}
			dp[i][1][j] = int(math.Max(
				float64(dp[i-1][1][j]),
				float64(dp[i-1][0][j]-prices[i]),
			))
		}
	}

	for j := 0; j < k+1; j++ {
		if res < dp[len(prices)-1][0][j] {
			res = dp[len(prices)-1][0][j]
		}
	}
	return res
}
