/*
https://leetcode-cn.com/contest/weekly-contest-235/problems/truncate-sentence/
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		s = "chopper is not a tanuki"
		k = 5
	)
	fmt.Printf("%s", truncateSentence(s, k))
}

func truncateSentence(s string, k int) string {
	var (
		sb    = strings.Builder{}
		count int
	)

	for i := range s {
		if s[i] == ' ' && (i == 0 || i == len(s)-1) {
			continue
		}
		if s[i] == ' ' {
			count++
			if count == k {
				break
			}
		}
		sb.WriteByte(s[i])
	}

	return sb.String()
}
