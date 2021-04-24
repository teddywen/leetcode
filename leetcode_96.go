/*
https://leetcode-cn.com/problems/unique-binary-search-trees/
 */
package main

import "fmt"

func main() {
	var (
		n = 18
	)
	fmt.Printf("%d", numTrees(n))
}

func numTrees(n int) int {
	if n == 0 {
		return 0
	}
	memo := make(map[int]map[int]int)
	return countNums(1, n, memo)
}

func countNums(start, end int, memo map[int]map[int]int) int {
	if start >= end {
		return 1
	}
	if res, ok := memo[start][end]; ok {
		return res
	}
	var res int
	for i:=start; i<=end; i++ {
		// left nums + right nums, if select i as the root
		res += countNums(start, i-1, memo) * countNums(i+1, end, memo)
	}
	if _, ok := memo[start]; !ok {
		memo[start] = make(map[int]int)
	}
	memo[start][end] = res
	return res
}
