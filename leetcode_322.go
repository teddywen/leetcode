/*
https://leetcode-cn.com/problems/coin-change/
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		coins  = []int{1}
		amount = 1
	)
	fmt.Printf("%d", coinChange(coins, amount))
}

func coinChange(coins []int, amount int) int {
	if amount == 0 || len(coins) == 0 {
		return 0
	}

	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		res := math.MaxInt32
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}
			if dp[i-coin] == -1 {
				continue // 无解
			}
			if dp[i-coin] + 1 < res {
				res = dp[i-coin] + 1
			}
		}
		if res == math.MaxInt32 {
			dp[i] = -1
		} else {
			dp[i] = res
		}
	}

	return dp[amount]
}
