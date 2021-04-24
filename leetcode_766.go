/*
https://leetcode-cn.com/problems/toeplitz-matrix/
*/
package main

import "fmt"

func main() {
	var (
		matrix = [][]int{{1, 2, 3, 4}, {5, 1, 2, 3}, {9, 5, 1, 2}}
	)
	fmt.Printf("%t", isToeplitzMatrix(matrix))
}

// 一次遍历法
// 沿上、左两条边开始作为起点，往右下方向做DFS。
func isToeplitzMatrix(matrix [][]int) bool {

	// j=0 遍历i
	for i := 0; i < len(matrix); i++ {
		if !_isToeplitzMatrix(i, 0, matrix) {
			return false
		}
	}

	// i=0 遍历j
	for j := 0; j < len(matrix[0]); j++ {
		if !_isToeplitzMatrix(0, j, matrix) {
			return false
		}
	}
	return true
}

func _isToeplitzMatrix(i, j int, matrix [][]int) bool {
	var val = -1
	for i < len(matrix) && j < len(matrix[0]) {
		if val == -1 {
			val = matrix[i][j]
		} else if val != matrix[i][j] {
			return false
		}
		i, j = i+1, j+1
	}
	return true
}

// 字典解法
//func isToeplitzMatrix(matrix [][]int) bool {
//	var (
//		dict = make(map[int]int)
//	)
//	for i := range matrix {
//		for j := range matrix[i] {
//			k := i-j
//			if v, ok := dict[k]; ok {
//				if v != matrix[i][j] {
//					return false
//				}
//			} else {
//				dict[k] = matrix[i][j]
//			}
//		}
//	}
//	return true
//}
