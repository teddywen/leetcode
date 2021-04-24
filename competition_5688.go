/*
https://leetcode-cn.com/contest/weekly-contest-229/problems/maximize-palindrome-length-from-subsequences/
*/
package main

import (
	"fmt"
)

func main() {
	var (
		word1 = "afaaadacb"
		word2 = "ca"
	)
	fmt.Printf("%d", longestPalindrome(word1, word2))
}

// 动规+路径压缩
func longestPalindrome(word1 string, word2 string) int {
	var (
		word = word1 + word2
		// 定义: dp[i][j][k] = x。
		// i,j: word字符串的起止指针。
		// x: 为在子串 word[i:j] 中最长回文子序列的长度
		// k: 0:最长回文子序列未跨过word1、word2。 1:最长回文子序列已跨过word1、word2
		// 求: dp[0][len(word)-1][1]
		// base-case: dp[i][i][0] = 1
		// -- 路径压缩说明 --
		// dp[i][j][k] 压缩投射至 dp[j][k]
		dp = make([][2]int, len(word))
	)

	for i := len(word) - 1; i >= 0; i-- {
		var temp0, temp1 int
		for j := i; j < len(word); j++ {
			if i == j {
				x := dp[j][0]
				dp[j][0] = 1
				temp0 = x
				continue
			}

			x, y := dp[j][0], dp[j][1]
			// i,j跨了左右串
			if i < len(word1) && j >= len(word1) {
				if word[i] == word[j] {
					dp[j][1] = max5688(temp0 + 2, temp1+2)
				} else {
					dp[j][1] = max5688(dp[j][1], dp[j-1][1])
					dp[j][0] = max5688(dp[j][0], dp[j-1][0])
				}
			} else {
				if word[i] == word[j] {
					dp[j][0] = temp0 + 2
				} else {
					dp[j][0] = max5688(dp[j][0], dp[j-1][0])
				}
			}
			temp0, temp1 = x, y
		}
	}

	return dp[len(word)-1][1]
}

//func longestPalindrome(word1 string, word2 string) int {
//	var (
//		word = word1 + word2
//		// 定义: dp[i][j][k] = x。
//		// i,j: word字符串的起止指针。
//		// x: 为在子串 word[i:j] 中最长回文子序列的长度
//		// k: 0:最长回文子序列未跨过word1、word2。 1:最长回文子序列已跨过word1、word2
//		// 求: dp[0][len(word)-1][1]
//		// base-case: dp[i][i][0] = 1
//		dp = make([][][2]int, len(word))
//	)
//	for i := range dp {
//		dp[i] = make([][2]int, len(word))
//	}
//
//	for i := len(word) - 1; i >= 0; i-- {
//		for j := i; j < len(word); j++ {
//			if i == j {
//				dp[i][j][0] = 1
//				continue
//			}
//
//			// i,j跨了左右串
//			if i < len(word1) && j >= len(word1) {
//				if word[i] == word[j] {
//					dp[i][j][1] = max5688(dp[i+1][j-1][0], dp[i+1][j-1][1]+2)
//				} else {
//					dp[i][j][1] = max5688(dp[i+1][j][1], dp[i][j-1][1])
//					dp[i][j][0] = max5688(dp[i+1][j][0], dp[i][j-1][0])
//				}
//			} else {
//				if word[i] == word[j] {
//					dp[i][j][0] = dp[i+1][j-1][0] + 2
//				} else {
//					dp[i][j][0] = max5688(dp[i+1][j][0], dp[i][j-1][0])
//				}
//			}
//		}
//	}
//
//	return dp[0][len(word)-1][1]
//}

func max5688(xyz ...int) int {
	var (
		v int
	)
	for i := range xyz {
		if xyz[i] > v {
			v = xyz[i]
		}
	}
	return v
}
