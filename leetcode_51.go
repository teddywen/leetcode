/*
https://leetcode-cn.com/problems/n-queens/
*/
package main

import "fmt"

func main() {
	var (
		n = 7
	)
	fmt.Printf("%#v", solveNQueens(n))
}

func solveNQueens(n int) [][]string {
	if n <= 0 {
		return [][]string{{}}
	}

	var (
		choice = n
		path   = make([]string, 0)
		chosen = make([]int, 0)
		result = make([][]string, 0)
	)

	backtraceQueue(&result, choice, chosen, path)

	return result
}

func backtraceQueue(result *[][]string, choice int, chosen []int, path []string) {
	if choice == 0 {
		var copyOf = make([]string, len(path))
		copy(copyOf, path)
		*result = append(*result, copyOf)
		return
	}

	var (
		n = len(chosen) + choice
		x = len(chosen)
	)

	for y := 0; y < n; y++ {

		pt := []int{x, y}

		// 做选择
		if !validQueue(chosen, pt) {
			continue
		}
		choice--
		exp := ""
		for i:=0; i<n; i++ {
			if i == y {
				exp += "Q"
			} else {
				exp += "."
			}
		}

		// 递归回溯
		backtraceQueue(result, choice, append(chosen, y), append(path, exp))

		// 撤销选择
		choice++
	}

}

func validQueue(chosen []int, pt []int) bool {
	for x, y := range chosen {
		if y == pt[1] {
			return false
		}
		if pt[0] - pt[1] == x - y {
			return false
		}
		if pt[0] + pt[1] == x + y {
			return false
		}
	}
	return true
}
