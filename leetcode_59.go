/*
https://leetcode-cn.com/problems/spiral-matrix-ii/
*/
package main

import "fmt"

func main() {
	var (
		n = 3
	)
	fmt.Printf("%#v", generateMatrix(n))
}

// 方向扫描法
//func generateMatrix(n int) [][]int {
//	var (
//		result            = make([][]int, n)
//		directs           = [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
//		i, j, directIndex int
//		count             = n * n
//	)
//	for i := range result {
//		result[i] = make([]int, n)
//	}
//
//	for k := 1; k <= count; k++ {
//		result[i][j] = k
//		nextI, nextJ := i+directs[directIndex][0], j+directs[directIndex][1]
//		if nextI < 0 || nextI > n-1 || nextJ < 0 || nextJ > n-1 || result[nextI][nextJ] > 0 {
//			directIndex = (directIndex + 1) % 4
//			nextI, nextJ = i+directs[directIndex][0], j+directs[directIndex][1]
//		}
//		i, j = nextI, nextJ
//	}
//
//	return result
//}

// 逐层遍历
func generateMatrix(n int) [][]int {
	var (
		result        = make([][]int, n)
		top, left     int
		right, bottom = n - 1, n - 1
		k             = 1
	)
	for i := range result {
		result[i] = make([]int, n)
	}

	for left <= right && top <= bottom {

		for i, j := top, left; j <= right; j++ {
			result[i][j] = k
			k++
		}

		for i, j := top+1, right; i <= bottom; i++ {
			result[i][j] = k
			k++
		}

		if left < right && top < bottom {

			for i, j := bottom, right-1; j >= left; j-- {
				result[i][j] = k
				k++
			}

			for i, j := bottom-1, left; i > top; i-- {
				result[i][j] = k
				k++
			}
		}

		left++
		right--
		top++
		bottom--
	}

	return result
}
