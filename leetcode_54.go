/*
https://leetcode-cn.com/problems/spiral-matrix/
*/
package main

import "fmt"

func main() {
	var (
		matrix = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	)
	fmt.Printf("%#v", spiralOrder(matrix))
}

// 方向扫描法
//func spiralOrder(matrix [][]int) []int {
//	var (
//		directs           = [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
//		i, j, directIndex int
//		n                 = len(matrix) * len(matrix[0])
//		result            []int
//	)
//
//	for k := 0; k < n; k++ {
//		result = append(result, matrix[i][j])
//		matrix[i][j] = 101 // cause -100 <= matrix[i][j] <= 100
//		nextI, nextJ := i + directs[directIndex][0], j + directs[directIndex][1]
//		if nextI < 0 || nextI > len(matrix)-1 || nextJ < 0 || nextJ > len(matrix[0])-1 || matrix[nextI][nextJ] == 101 {
//			directIndex = (directIndex + 1) % 4
//			nextI, nextJ = i + directs[directIndex][0], j + directs[directIndex][1]
//		}
//		i, j = nextI, nextJ
//	}
//
//	return result
//}

// 逐层遍历
func spiralOrder(matrix [][]int) []int {
	var (
		leftTop     = [2]int{0, 0}
		rightBottom = [2]int{len(matrix) - 1, len(matrix[0]) - 1}
		result      []int
	)

	for leftTop[0] <= rightBottom[0] && leftTop[1] <= rightBottom[1] {

		for i, j := leftTop[0], leftTop[1]; j <= rightBottom[1]; j++ {
			result = append(result, matrix[i][j])
		}

		for i, j := leftTop[0]+1, rightBottom[1]; i <= rightBottom[0]; i++ {
			result = append(result, matrix[i][j])
		}

		if leftTop[0] < rightBottom[0] && leftTop[1] < rightBottom[1] {

			for i, j := rightBottom[0], rightBottom[1]-1; j >= leftTop[1]; j-- {
				result = append(result, matrix[i][j])
			}

			for i, j := rightBottom[0]-1, leftTop[1]; i > leftTop[0]; i-- {
				result = append(result, matrix[i][j])
			}
		}

		leftTop[0]++
		leftTop[1]++
		rightBottom[0]--
		rightBottom[1]--
	}

	return result
}
