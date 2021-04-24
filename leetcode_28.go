/*
https://leetcode-cn.com/problems/implement-strstr/
 */
package main

import "fmt"

func main() {
	var (
		haystack = "abc"
		needle = "c"
	)
	fmt.Printf("%d", strStr(haystack, needle))
}

func strStr(haystack string, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	}

	var (
		end = len(needle) - 1
	)

	for end < len(haystack) {

		end++

		start := end-len(needle)
		if haystack[start:end] == needle {
			return start
		}

	}

	return -1
}
