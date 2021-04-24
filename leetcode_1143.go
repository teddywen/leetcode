/*
https://leetcode-cn.com/problems/longest-common-subsequence/
*/
package main

import (
	"fmt"
)

func main() {
	var (
		text1 = "abcde"
		text2 = "ace"
	)
	fmt.Printf("%d", longestCommonSubsequence(text1, text2))
}

// 路径压缩
func longestCommonSubsequence(text1 string, text2 string) int {
	var (
		// 定义: dp[i][j] = x。text1[1~i]和text2[1~j]的最长公共子序列的长度为x
		// base case: dp[0][?] = 0, dp[?][0] = 0
		// 求: dp[len(text1)][len(text2)]
		// dp方程:
		// text1[i] == text2[j]: dp[i][j] = dp[i-1][j-1] + 1
		// else: dp[i][j] = max(dp[i-1][j], dp[i][j-1])
		// 在dp[m][n]的矩阵中，由于dp[i][j]只依赖上、左、左上的数据。因此可以把i压缩掉，dp投射至一维数组dp[j]。
		dp = make([]int, len(text2)+1)
	)
	for i := 1; i <= len(text1); i++ {
		var temp int
		for j := 1; j <= len(text2); j++ {
			x := dp[j]
			if text1[i-1] == text2[j-1] {
				dp[j] = temp + 1
			} else {
				dp[j] = max1143(dp[j-1], dp[j])
			}
			temp = x
		}
	}
	return dp[len(text2)]
}

//func longestCommonSubsequence(text1 string, text2 string) int {
//	var (
//		// 定义: dp[i][j] = x。text1[1~i]和text2[1~j]的最长公共子序列的长度为x
//		// base case: dp[0][?] = 0, dp[?][0] = 0
//		// 求: dp[len(text1)][len(text2)]
//		// dp方程:
//		// text1[i] == text2[j]: dp[i][j] = dp[i-1][j-1] + 1
//		// else: dp[i][j] = max(dp[i-1][j], dp[i][j-1])
//		dp = make([][]int, len(text1)+1)
//	)
//	for i := range dp {
//		dp[i] = make([]int, len(text2)+1)
//	}
//	for i := 1; i <= len(text1); i++ {
//		for j := 1; j <= len(text2); j++ {
//			if text1[i-1] == text2[j-1] {
//				dp[i][j] = dp[i-1][j-1] + 1
//			} else {
//				dp[i][j] = max1143(dp[i][j-1], dp[i-1][j])
//			}
//		}
//	}
//	return dp[len(text1)][len(text2)]
//}

func max1143(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
