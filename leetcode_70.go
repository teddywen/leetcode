/*
https://leetcode-cn.com/problems/climbing-stairs/
 */
package main

import "fmt"

func main() {
	var (
		n = 1
	)
	fmt.Printf("%d", climbStairs(n))
}

func climbStairs(n int) int {
	var dp = [2]int{0, 0}
	for i:=0; i<=n; i++ {
		if i < 2 {
			dp[0], dp[1] = dp[1], 1
		} else {
			dp[0], dp[1] = dp[1], dp[0] + dp[1]
		}
	}
	return dp[1]
}

//func climbStairs(n int) int {
//	var dp = make([]int, n+1)
//	for i:=0; i<=n; i++ {
//		if i < 2 {
//			dp[i] = 1
//		} else {
//			dp[i] = dp[i-1] + dp[i-2]
//		}
//	}
//	return dp[n]
//}

//// 穷举
//func climbStairs(n int) int {
//	return climbStairsBacktrace(0, n)
//}
//
//// 多少种走法
//func climbStairsBacktrace(i, n int) int {
//	if i == n {
//		return 1
//	}
//	if i > n {
//		return 0
//	}
//
//	// 选择走1格
//	i += 1
//	choice1 := climbStairsBacktrace(i, n)
//	// 撤销
//	i -= 1
//
//	// 选择走2格
//	i += 2
//	choice2 := climbStairsBacktrace(i, n)
//	i -= 2
//
//	return choice1 + choice2
//}
