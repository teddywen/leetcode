/*
https://leetcode-cn.com/problems/pascals-triangle-ii/
*/
package main

import "fmt"

func main() {
	var (
		rowIndex = 3
	)
	fmt.Printf("%#v", getRow(rowIndex))
}

func getRow(rowIndex int) []int {
	var (
		// dp数组二维变一维，空间压缩
		dp = make([]int, rowIndex+1)
	)
	dp[0] = 1

	// 从后往前可以让dp一维数组信息不被覆盖，避免申请临时变量
	for i := 0; i < len(dp); i++ {
		for j := i; j > 0; j-- {
			dp[j] = dp[j] + dp[j-1]
		}
	}
	return dp
}
