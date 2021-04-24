/*
https://leetcode-cn.com/problems/palindrome-partitioning-ii/
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		s = "aab"
	)
	fmt.Printf("%d", minCut(s))
}

func minCut(s string) int {
	var (
		isPalindrome = make([][]bool, len(s))
		// 定义: dp[i] = x
		// 字符串第0~i的最小分割次数
		// 方程: dp[i] = dp[j] + 1
		// 当字符串j+1 ~ n-1为回文串时，dp[i] = dp[j] + 1
		dp = make([]int, len(s))
	)
	// 初始化dp
	for i := 0; i < len(s); i++ {
		isPalindrome[i] = make([]bool, len(s))
		for j := 0; j < len(s); j++ {
			isPalindrome[i][j] = true
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		for j := i + 1; j < len(s); j++ {
			isPalindrome[i][j] = s[i] == s[j] && isPalindrome[i+1][j-1]
		}
	}
	for i := 0; i < len(s); i++ {
		// 当字符串0~i本身是个回文串时，最小切割次数为0
		if isPalindrome[0][i] {
			continue
		}
		dp[i] = math.MaxInt32
		for j := i - 1; j >= 0; j-- {
			if isPalindrome[j+1][i] {
				dp[i] = min132(dp[i], dp[j]+1)
			}
		}
	}

	return dp[len(s)-1]
}

func min132(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
