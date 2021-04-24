/*
https://leetcode-cn.com/problems/minimum-insertion-steps-to-make-a-string-palindrome/
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		s = "leetcode"
	)
	fmt.Printf("%d", minInsertions(s))
}

// 路径压缩版
func minInsertions(s string) int {
	var (
		dp = make([]int, len(s))
	)

	// dp[i][j] = x
	// 定义为 s[i:j+1]变成回文串的最小插入次数
	for i := len(s) - 2; i >= 0; i-- {
		// 记录每次里层循环的每次迭代的上一次迭代未被覆盖时的状态值
		var pre int
		for j := i + 1; j < len(s); j++ {
			// 再即将被覆盖前复制出来
			tmp := dp[j]

			if s[i] == s[j] {
				dp[j] = pre
			} else {
				dp[j] = int(math.Min(float64(dp[j-1]), float64(dp[j]))) + 1
			}

			// 操作完在保存到pre变量，这样下次迭代时pre的值就是上一次未覆盖时的值了
			pre = tmp
		}
	}

	return dp[len(s)-1]
}


//func minInsertions(s string) int {
//	var (
//		dp = make([][]int, len(s))
//	)
//	for i := 0; i < len(s); i++ {
//		dp[i] = make([]int, len(s))
//	}
//
//	// dp[i][j] = x
//	// 定义为 s[i:j+1]变成回文串的最小插入次数
//	for i := len(s) - 2; i >= 0; i-- {
//		for j := i + 1; j < len(s); j++ {
//			if s[i] == s[j] {
//				dp[i][j] = dp[i+1][j-1]
//			} else {
//				dp[i][j] = int(math.Min(float64(dp[i][j-1]), float64(dp[i+1][j]))) + 1
//			}
//		}
//	}
//
//	return dp[0][len(s)-1]
//}
