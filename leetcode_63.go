/*
https://leetcode-cn.com/problems/unique-paths-ii/
 */
package main

import "fmt"

func main() {
	var (
		obstacleGrid = [][]int{{0,0,0},{0,1,0},{0,0,0}}
	)
	fmt.Printf("%d", uniquePathsWithObstacles(obstacleGrid))
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	// dp[i][j] = x 从起始到坐标(i,j)一共有几种走法
	var (
		cols = len(obstacleGrid[0])
		rows = len(obstacleGrid)
		//dp [][]int
	)

	//dp = make([][]int, cols)
	//for i := range dp {
	//	dp[i] = make([]int, rows)
	//}
	//dp[0][0] = 1

	if !valid(cols-1, rows-1, obstacleGrid) {
		return 0
	}

	// dp[i][j] needs dp[i-1][j] dp[i][j-1]
	dp := [2][2]int{}

	// dp[0][0] -> dp[cols-1][rows-1]
	for i:=0; i<cols; i++ {
		for j:=0; j<rows; j++ {
			if i == 0 && j == 0 {
				zoo = 1
				//dp[i][j] = 1
			} else {
				var sum int
				// up
				if valid(i, j-1, obstacleGrid) {
					sum += bar
					//sum += dp[i][j-1]
				}
				// left
				if valid(i-1, j, obstacleGrid) {
					sum += foo
					//sum += dp[i-1][j]
				}
				dp[i][j] = sum
			}
		}
		dp[0] = dp[1]
	}

	return dp[cols-1][rows-1]
}

//func uniquePathsWithObstacles(obstacleGrid [][]int) int {
//	// dp[i][j] = x 从起始到坐标(i,j)一共有几种走法
//	var (
//		cols = len(obstacleGrid[0])
//		rows = len(obstacleGrid)
//		dp [][]int
//	)
//
//	dp = make([][]int, cols)
//	for i := range dp {
//		dp[i] = make([]int, rows)
//	}
//	dp[0][0] = 1
//
//	if !valid(cols-1, rows-1, obstacleGrid) {
//		return 0
//	}
//
//	// dp[0][0] -> dp[cols-1][rows-1]
//	for i:=0; i<cols; i++ {
//		for j:=0; j<rows; j++ {
//			if i == 0 && j == 0 {
//				dp[i][j] = 1
//			} else {
//				var sum int
//				// up
//				if valid(i, j-1, obstacleGrid) {
//					sum += dp[i][j-1]
//				}
//				// left
//				if valid(i-1, j, obstacleGrid) {
//					sum += dp[i-1][j]
//				}
//				dp[i][j] = sum
//			}
//		}
//	}
//
//	return dp[cols-1][rows-1]
//}

func valid(x, y int, obstacleGrid [][]int) bool {
	if x < 0 || y < 0 {
		return false
	}
	if obstacleGrid[y][x] == 1 {
		return false
	}
	return true
}