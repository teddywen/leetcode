/*
https://leetcode-cn.com/problems/range-sum-query-2d-immutable/
*/
package main

import "fmt"

func main() {
	var (
		matrix = [][]int{
			{3, 0, 1, 4, 2},
			{5, 6, 3, 2, 1},
			{1, 2, 0, 1, 5},
			{4, 1, 0, 1, 7},
			{1, 0, 3, 0, 5},
		}
		obj = Constructor(matrix)
	)
	fmt.Printf("%#v", obj.SumRegion(2, 1, 4, 3))
}

type NumMatrix struct {
	prefix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	var (
		dp = make([][]int, len(matrix)+1)
	)
	if len(matrix) > 0 {
		for i := range dp {
			dp[i] = make([]int, len(matrix[0])+1)
		}
	}
	// dp[i][j] = x
	// dp[i][j] = dp[i-1][j] + dp[i][j-1] - dp[i-1][j-1] + matrix[i][j]
	for i := 1; i < len(matrix)+1; i++ {
		for j := 1; j < len(matrix[0])+1; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1] - dp[i-1][j-1] + matrix[i-1][j-1]
		}
	}
	return NumMatrix{prefix: dp}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.prefix[row2+1][col2+1] - this.prefix[row1][col2+1] - this.prefix[row2+1][col1] + this.prefix[row1][col1]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
