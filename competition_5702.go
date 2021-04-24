/*
https://leetcode-cn.com/contest/weekly-contest-232/problems/find-center-of-star-graph/
*/
package main

import (
	"fmt"
)

func main() {
	var (
		edges = [][]int{{1,2},{5,1},{1,3},{1,4}}
	)
	fmt.Printf("%d", findCenter(edges))
}

func findCenter(edges [][]int) int {
	if edges[0][0] == edges[1][0] {
		return edges[0][0]
	} else if edges[0][0] == edges[1][1] {
		return edges[0][0]
	} else if edges[0][1] == edges[1][0] {
		return edges[0][1]
	} else {
		return edges[0][1]
	}
}