/*
https://leetcode-cn.com/problems/transpose-matrix/
*/
package main

import "fmt"

func main() {
	var (
		matrix = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
		//matrix = [][]int{{1, 2, 3}, {4, 5, 6}}
	)
	fmt.Printf("%#v", transpose(matrix))
}

func transpose(matrix [][]int) [][]int {
	var (
		reverse = make([][]int, len(matrix[0]))
	)
	for i := range reverse {
		reverse[i] = make([]int, len(matrix))
	}
	for i := range matrix {
		for j := range matrix[i] {
			reverse[j][i] = matrix[i][j]
		}
	}
	return reverse
}
