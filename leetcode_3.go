/*
https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
 */
package main

import "fmt"

func main() {
	var (
		s = "abba"
	)
	fmt.Printf("%#v", lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {
	var (
		left, right, res int
		dict  = make(map[byte]int)
	)
	for right < len(s) {
		c := s[right]
		right++

		if to, ok := dict[c]; ok && to > left {
			left = to
		}

		res = max3(res, right - left)
		dict[c] = right
	}
	return res
}

func max3(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}