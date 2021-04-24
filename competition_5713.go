/*
https://leetcode-cn.com/problems/number-of-different-integers-in-a-string/
*/
package main

import (
	"fmt"
)

func main() {
	var (
		word = "a1b01c001"
	)
	fmt.Printf("%d", numDifferentIntegers(word))
}

func numDifferentIntegers(word string) int {
	var (
		count int
		sum   = -1
		memo  = make(map[int]struct{})
	)

	for i := range word {
		if word[i] >= '0' && word[i] <= '9' {
			if sum == -1 {
				sum = int(word[i] - '0')
			} else {
				sum = int(word[i]-'0') + sum*10
			}
		} else {
			if sum != -1 {
				if _, ok := memo[sum]; !ok {
					count++
					memo[sum] = struct{}{}
				}
				sum = -1
			}
		}
	}
	if sum != -1 {
		if _, ok := memo[sum]; !ok {
			count++
			memo[sum] = struct{}{}
		}
	}

	return count
}
