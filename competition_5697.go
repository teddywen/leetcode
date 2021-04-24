/*
https://leetcode-cn.com/contest/weekly-contest-231/problems/check-if-binary-string-has-at-most-one-segment-of-ones/
*/
package main

import (
	"fmt"
)

func main() {
	var (
		s = "1101"
	)
	fmt.Printf("%t", checkOnesSegment(s))
}

func checkOnesSegment(s string) bool {
	var (
		firstOne = -1
		lastOne = -1
	)

	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			firstOne = i
			break
		}
	}
	if firstOne == -1 {
		return false
	}

	for j:=len(s)-1; j>=0; j-- {
		if s[j] == '1' {
			lastOne = j
			break
		}
	}

	for i:=firstOne+1; i<lastOne; i++ {
		if s[i] == '0' {
			return false
		}
	}

	return true
}
