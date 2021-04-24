/*
https://leetcode-cn.com/problems/set-matrix-zeroes/
*/
package main

import "fmt"

func main() {
	var (
		matrix = [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	)
	setZeroes(matrix)
	fmt.Printf("%v", matrix)
}

func setZeroes(matrix [][]int) {
	var (
		colZero bool
	)
	for i := range matrix {
		for j, v := range matrix[i] {
			if v != 0 {
				continue
			}
			// 记录行为0
			matrix[i][0] = 0
			// 记录列为0
			if j == 0 {
				colZero = true
			} else {
				matrix[0][j] = 0
			}
		}
	}

	for i := len(matrix) - 1; i >= 0; i-- {
		for j := len(matrix[0]) - 1; j > 0; j-- {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
		if colZero {
			matrix[i][0] = 0
		}
	}
}
