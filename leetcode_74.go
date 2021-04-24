/*
https://leetcode-cn.com/problems/two-sum/
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		matrix = [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
		target = 30
	)
	fmt.Printf("%t", searchMatrix(matrix, target))
}

func searchMatrix(matrix [][]int, target int) bool {
	var (
		left  = 0
		right = len(matrix)*len(matrix[0]) - 1
	)
	for left <= right {
		mid := left + (right-left)/2
		val := matrixVal(matrix, mid)
		if target == val {
			return true
		} else if target < val {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}

func matrixVal(matrix [][]int, idx int) int {
	i := int(math.Floor(float64(idx / len(matrix[0]))))
	j := idx % len(matrix[0])
	return matrix[i][j]
}
