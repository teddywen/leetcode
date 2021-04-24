/*
https://leetcode-cn.com/problems/distinct-subsequences/
*/
package main

import "fmt"

func main() {
	var (
		s = "babgbag"
		t = "bag"
	)
	fmt.Printf("%d", numDistinct(s, t))
}

// 空间压缩
func numDistinct(s, t string) int {
	if len(s) < len(t) {
		return 0
	}

	var (
		// dp定义: dp[i][j] = x
		// x为"s串的前i个子串"和"t串的前j个子串"的子序列个数
		// dp状态转移方程:
		// if s[i-1]==[j-1]
		//			dp[i][j] = dp[i-1][j-1] + dp[i-1][j]  // 末尾包含s[i]的子序列数 + 末尾不包含s[i]的子序列数
		// else
		//			dp[i][j] = dp[i-1][j]	// 末尾不包含s[i]的子序列数。此情况下包含后s[i]后不构成子序列
		// base case:
		// dp[0][0] = 1, dp[i][0] = 1, dp[0][j] = 0		// 根据基础状况反推出，脑内模拟当 s=ab,t=b 时base case的值如何表现
		// 精简为 dp[i][0] = 1，其余为0
		// 求 dp[i][j] i=len(s), j=len(t)
		dp = make([]int, len(t)+1)
	)
	for i := 0; i <= len(s); i++ {
		var pre int
		for j := 0; j <= len(t); j++ {
			tmp := dp[j]
			if j == 0 {
				dp[j] = 1
			} else if i == 0 {
				dp[j] = 0
			} else if s[i-1] == t[j-1] {
				dp[j] = pre + dp[j]
			} else {
				dp[j] = dp[j]
			}
			pre = tmp
		}
	}
	return dp[len(t)]
}

//func numDistinct(s, t string) int {
//	if len(s) < len(t) {
//		return 0
//	}
//
//	var (
//		// dp定义: dp[i][j] = x
//		// x为"s串的前i个子串"和"t串的前j个子串"的子序列个数
//		// dp状态转移方程:
//		// if s[i-1]==[j-1]
//		//			dp[i][j] = dp[i-1][j-1] + dp[i-1][j]  // 末尾包含s[i]的子序列数 + 末尾不包含s[i]的子序列数
//		// else
//		//			dp[i][j] = dp[i-1][j]	// 末尾不包含s[i]的子序列数。此情况下包含后s[i]后不构成子序列
//		// base case:
//		// dp[0][0] = 1, dp[i][0] = 1, dp[0][j] = 0		// 根据基础状况反推出，脑内模拟当 s=ab,t=b 时base case的值如何表现
//		// 精简为 dp[i][0] = 1，其余为0
//		// 求 dp[i][j] i=len(s), j=len(t)
//		dp = make([][]int, len(s)+1)
//	)
//	for i := range dp {
//		dp[i] = make([]int, len(t)+1)
//		dp[i][0] = 1
//	}
//	for i := 1; i <= len(s); i++ {
//		for j := 1; j <= len(t); j++ {
//			if s[i-1] == t[j-1] {
//				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
//			} else {
//				dp[i][j] = dp[i-1][j]
//			}
//		}
//	}
//	return dp[len(s)][len(t)]
//}
