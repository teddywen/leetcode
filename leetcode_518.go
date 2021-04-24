/*
https://leetcode-cn.com/problems/coin-change-2/
*/
package main

import "fmt"

func main() {
	var (
		amount = 5
		coins  = []int{1, 2, 5}
	)
	fmt.Printf("%d", change(amount, coins))
}

//func change(amount int, coins []int) int {
//	if amount <= 0 {
//		return 1
//	}
//	if len(coins) == 0 {
//		return 0
//	}
//
//	var (
//		dp = make([][]int, len(coins)+1)
//	)
//	// 当要凑出的金额为0时，不凑也是一种凑法
//	for i := range dp {
//		dp[i] = make([]int, amount+1)
//		dp[i][0] = 1
//	}
//	// 当手里没有硬币时，任何金额都无法凑出。
//	for i := range dp[0] {
//		dp[0][i] = 0
//	}
//
//	// dp[i][j] 凑硬币的组合数 i:第i枚硬币,i=0代表不选 j:剩余amount
//	for i := 1; i <= len(coins); i++ {
//		for j := 1; j <= amount; j++ {
//			if j >= coins[i-1] {
//				// 一种是什么都不选，因此组合数即为dp[i-1][j]
//				// 一种是选一枚当前的硬币，因此组合数为dp[i][j-coins[i-1]]
//				// 两种选择加起来就是所有组合数
//				dp[i][j] = dp[i-1][j] + dp[i][j-coins[i-1]]
//			} else {
//				dp[i][j] = dp[i-1][j]
//			}
//		}
//	}
//
//	return dp[len(coins)][amount]
//}

// 空间压缩
func change(amount int, coins []int) int {
	// 当要凑出的金额为0时，不凑也是一种凑法
	if amount <= 0 {
		return 1
	}
	// 当手里没有硬币时，任何金额都无法凑出。
	if len(coins) == 0 {
		return 0
	}

	var (
		dp = make([]int, amount+1)
	)
	dp[0] = 1

	// dp[i][j] 凑硬币的组合数 i:第i枚硬币,i=0代表不选 j:剩余amount
	for i := 1; i <= len(coins); i++ {
		for j := 1; j <= amount; j++ {
			if j >= coins[i-1] {
				// 一种是什么都不选，因此组合数即为dp[i-1][j]
				// 一种是选一枚当前的硬币，因此组合数为dp[i][j-coins[i-1]]
				// 两种选择加起来就是所有组合数
				last := dp
				dp[j] = last[j] + dp[j-coins[i-1]]
			} else {
				last := dp
				dp[j] = last[j]
			}
		}
	}

	return dp[amount]
}
