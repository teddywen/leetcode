/*
https://leetcode-cn.com/problems/longest-palindromic-substring/
*/
package main

import "fmt"

func main() {
	var (
		s = "abcba"
	)
	fmt.Printf("%s\n", longestPalindrome(s))
}

// 动态规划 + 状态压缩
func longestPalindrome(s string) string {
	var (
		// dp定义: dp[i][j] = x
		// x为s[i:j]是否为回文子串
		// dp方程:
		// 		if s[i] == s[j] && dp[i-1][j-1]: dp[i][j] = dp[i-1][j-1]
		// base case: dp[i][i] = true
		dp = make([]bool, len(s))
		res = s[:1]
	)

	for i := len(s) - 2; i >= 0; i-- {
		pre := false
		for j := i + 1; j < len(s); j++ {

			tmp := dp[j]

			if pre {
				dp[j] = true
			} else {
				if s[i] == s[j] {
					dp[j] = false
					if j - i + 1 > len(res) {
						res = s[i:j+1]
					}
				} else {
					dp[j] = true
				}
			}

			pre = tmp
		}
	}
	return res
}


// 动态规划
//func longestPalindrome(s string) string {
//	var (
//		// dp定义: dp[i][j] = x
//		// x为s[i:j]是否为回文子串
//		// dp方程:
//		// 		if s[i] == s[j] && dp[i-1][j-1]: dp[i][j] = dp[i-1][j-1]
//		// base case: dp[i][i] = true
//		dp = make([][]bool, len(s))
//		res = s[:1]
//	)
//	for i := range dp {
//		dp[i] = make([]bool, len(s))
//		for j := range dp[i] {
//			dp[i][j] = true
//		}
//	}
//
//	for i := len(s) - 2; i >= 0; i-- {
//		for j := i + 1; j < len(s); j++ {
//			if !dp[i+1][j-1] {
//				dp[i][j] = false
//				continue
//			}
//			if s[i] == s[j] {
//				dp[i][j] = true
//				if j - i + 1 > len(res) {
//					res = s[i:j+1]
//				}
//			} else {
//				dp[i][j] = false
//			}
//		}
//	}
//	return res
//}
