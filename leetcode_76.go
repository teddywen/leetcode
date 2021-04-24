/*
https://leetcode-cn.com/problems/minimum-window-substring/
*/
package main

import "fmt"

func main() {
	var (
		s = "ADOBECODEBANC"
		t = "ABC"
	)
	fmt.Printf("%s\n", minWindow(s, t))
}

func minWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 || len(s) < len(t) {
		return ""
	}

	var (
		need   = make(map[uint8]int)
		window = make(map[uint8]int)
		res = ""
	)
	var left, right, valid int

	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	for right < len(s) {
		// expand
		c := s[right]
		right++

		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		for valid == len(need) {
			if valid == len(need) && (len(res) > right - left || res == ""){
				res = s[left:right]
			}

			// shrink
			c := s[left]
			left++

			if window[c] > 0 {
				if window[c] == need[c] {
					valid--
				}
				window[c]--
				if window[c] == 0 {
					delete(window, c)
				}
			}
		}
	}

	return res
}
