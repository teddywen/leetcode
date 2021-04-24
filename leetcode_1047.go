/*
https://leetcode-cn.com/problems/remove-all-adjacent-duplicates-in-string/
*/
package main

import (
	"fmt"
)

func main() {
	var (
		S = "abbaca"
	)
	fmt.Printf("%s", removeDuplicates(S))
}

func removeDuplicates(S string) string {
	var (
		buf []byte
	)
	for i := range S {
		if len(buf) == 0 || buf[len(buf)-1] != S[i] {
			buf = append(buf, S[i])
		} else {
			buf = buf[:len(buf)-1]
		}
	}
	return string(buf)
}
