/*
https://leetcode-cn.com/contest/weekly-contest-229/problems/minimum-number-of-operations-to-move-all-balls-to-each-box/
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		boxes = "110"
	)
	fmt.Printf("%#v", minOperations(boxes))
}

// O(n2)
func minOperations(boxes string) []int {
	var (
		idx = make([]int, len(boxes))
		res = make([]int, len(boxes))
	)
	for i := range boxes {
		if boxes[i] == '1' {
			idx[i] = i
		} else {
			idx[i] = -1
		}
	}
	for i := 0; i < len(boxes); i++ {
		for j := 0; j < len(boxes); j++ {
			if i == j {
				continue
			}
			res[i] += abs5686(idx[j], i)
		}
	}
	return res
}

func abs5686(x, y int) int {
	if x == -1 {
		return 0
	}
	return int(math.Abs(float64(x - y)))
}
