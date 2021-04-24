/*
https://leetcode-cn.com/problems/edit-distance/
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		word1 = "intention"
		word2 = "execution"
	)
	fmt.Printf("%d", minDistance(word1, word2))
}

// 路径压缩
func minDistance(word1 string, word2 string) int {
	var (
		// 定义: dp[i][j] = x。word1[1~i]和word2[1~j]的最少编辑次数
		// base case: dp[0][j] = j, dp[i][0] = i
		// 求: dp[len(word1)][len(word2)]
		// dp方程:
		// word1[i] == word2[j]: dp[i][j] = dp[i-1][j-1]
		// else: dp[i][j] = min(
		//		dp[i][j-1] + 1, 	// 增
		//		dp[i-1][j] + 1, 	// 删
		//		dp[i-1][j-1] + 1,	// 改
		// )
		// 在dp[m][n]的矩阵中，由于dp[i][j]只依赖上、左、左上的数据。因此可以把i压缩掉，dp投射至一维数组dp[j]。
		dp = make([]int, len(word2)+1)
	)

	for i := 0; i <= len(word1); i++ {
		var temp int
		for j := 0; j <= len(word2); j++ {
			// base case
			if i == 0 {
				dp[j] = j
				continue
			}
			if j == 0 {
				x := dp[j]
				dp[j] = i
				temp = x
				continue
			}

			x := dp[j]
			if word1[i-1] == word2[j-1] {
				dp[j] = temp
			} else {
				dp[j] = min72(dp[j-1]+1, dp[j]+1, temp+1)
			}
			temp = x
		}
	}

	return dp[len(word2)]
}

//func minDistance(word1 string, word2 string) int {
//	var (
//		// 定义: dp[i][j] = x。word1[1~i]和word2[1~j]的最少编辑次数
//		// base case: dp[0][j] = j, dp[i][0] = i
//		// 求: dp[len(word1)][len(word2)]
//		// dp方程:
//		// word1[i] == word2[j]: dp[i][j] = dp[i-1][j-1]
//		// else: dp[i][j] = min(
//		//		dp[i][j-1] + 1, 	// 增
//		//		dp[i-1][j] + 1, 	// 删
//		//		dp[i-1][j-1] + 1,	// 改
//		// )
//		dp = make([][]int, len(word1)+1)
//	)
//	for i := range dp {
//		dp[i] = make([]int, len(word2)+1)
//	}
//
//	for i := 0; i <= len(word1); i++ {
//		for j := 0; j <= len(word2); j++ {
//			// base case
//			if i == 0 || j == 0 {
//				dp[i][j] = i + j
//				continue
//			}
//
//			if word1[i-1] == word2[j-1] {
//				dp[i][j] = dp[i-1][j-1]
//			} else {
//				dp[i][j] = min72(dp[i][j-1]+1, dp[i-1][j]+1, dp[i-1][j-1]+1)
//			}
//		}
//	}
//
//	return dp[len(word1)][len(word2)]
//}

func min72(xyz ...int) int {
	var (
		res = math.MaxInt32
	)
	for _, v := range xyz {
		if v < res {
			res = v
		}
	}
	return res
}
