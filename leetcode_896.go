/*
https://leetcode-cn.com/problems/monotonic-array/
*/
package main

import "fmt"

func main() {
	var (
		A = []int{6, 5, 4, 4}
	)
	fmt.Printf("%#v", isMonotonic(A))
}

func isMonotonic(A []int) bool {
	var (
		lastCmpVal = 0
	)
	for i := 0; i < len(A)-1; i++ {
		if A[i] != A[i+1] {
			cmpVal := compare896(A[i], A[i+1])
			if lastCmpVal != 0 && lastCmpVal != cmpVal {
				return false
			}
			lastCmpVal = cmpVal
		}
	}
	return true
}

func compare896(x, y int) int {
	if x < y {
		return -1
	} else {
		return 1
	}
}
