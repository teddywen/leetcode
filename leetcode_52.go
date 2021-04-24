/*
https://leetcode-cn.com/problems/n-queens-ii/
*/
package main

import (
	"fmt"
)

func main() {
	var (
		n = 4
	)
	fmt.Printf("%d", totalNQueens(n))
}

func totalNQueens(n int) int {
	return backtrace(-1, nil, nil, nil, n)
}

// pie: x+y
// na: |x-y|
func backtrace(row int, cols []int, pie []int, na []int, n int) int {
	if row == n-1 {
		return 1
	}

	var res int
	row += 1
	for i := 0; i < n; i++ {
		// 选择
		if !safe(row, i, cols, pie, na) {
			continue
		}
		cols = append(cols, i)
		pie = append(pie, row+i)
		na = append(na, row-i)
		// 递归
		res += backtrace(row, cols, pie, na, n)
		// 撤销
		cols = cols[:len(cols)-1]
		pie = pie[:len(pie)-1]
		na = na[:len(na)-1]
	}
	return res
}

func safe(x, y int, cols []int, pie []int, na []int) bool {
	for _, c := range cols {
		if y == c {
			return false
		}
	}
	for _, p := range pie {
		if x+y == p {
			return false
		}
	}
	for _, n := range na {
		if x-y == n {
			return false
		}
	}
	return true
}
