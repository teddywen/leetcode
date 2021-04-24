/*
https://leetcode-cn.com/contest/weekly-contest-232/problems/check-if-one-string-swap-can-make-strings-equal/
*/
package main

import (
	"fmt"
)

func main() {
	var (
		s1 = "ac"
		s2 = "bc"
	)
	fmt.Printf("%t", areAlmostEqual(s1, s2))
}

func areAlmostEqual(s1 string, s2 string) bool {
	var (
		diff   int
		wrong2 byte
		wrong1 byte
	)

	for i := range s2 {
		if s2[i] != s1[i] {

			if diff == 2 {
				return false
			}

			if diff == 0 {
				wrong2, wrong1 = s2[i], s1[i]
				diff++
			} else {
				if s2[i] != wrong1 || s1[i] != wrong2 {
					return false
				}
				diff++
			}
		}
	}
	return diff != 1
}