/*
https://leetcode-cn.com/contest/weekly-contest-231/problems/number-of-restricted-paths-from-first-to-last-node/
*/
package main

import (
	"fmt"
)

func main() {
	var (
		n     = 5
		edges = [][]int{{1, 2, 3}, {1, 3, 3}, {2, 3, 1}, {1, 4, 2}, {5, 2, 2}, {3, 5, 1}, {5, 4, 10}}
	)
	fmt.Printf("%d", countRestrictedPaths(n, edges))
}

func countRestrictedPaths(n int, edges [][]int) int {
	return 0
}
