/*
https://leetcode-cn.com/problems/longest-substring-with-at-least-k-repeating-characters/
*/
package main

import "fmt"

func main() {
	var (
		s = "aaabb"
		k = 3
		//s = "ababbc"
		//k = 2
		//s = "ababacb"
		//k = 3
	)
	fmt.Printf("%d", longestSubstring(s, k))
}

// 分治
func longestSubstring(s string, k int) int {
	return findLSS(0, len(s)-1, k, s)
}

func findLSS(left, right, k int, s string) int {
	if left > right {
		return 0
	}

	var cntDict = make(map[byte]int)

	for i := left; i <= right; i++ {
		cntDict[s[i]]++
	}

	for i := left; i <= right; i++  {
		if cntDict[s[i]] < k {
			return max395(findLSS(left, i-1, k, s), findLSS(i+1, right, k, s))
		}
	}
	return right - left + 1
}

// 滑动窗口
//func longestSubstring(s string, k int) int {
//	var res int
//
//	// 自己制造限制，让其满足滑窗条件
//	for t := 1; t <= 26; t++ {
//		var (
//			left, right, spec, bad int
//			cntDict                = make(map[byte]int)
//		)
//
//		for right < len(s) {
//			c := s[right]
//			right++
//
//			if cntDict[c] == 0 {
//				spec++
//				bad++
//			}
//			cntDict[c]++
//			if cntDict[c] == k {
//				bad--
//			}
//
//			for left < right && spec > t {
//				c2 := s[left]
//				left++
//
//				if cntDict[c2] == k {
//					bad++
//				}
//				if cntDict[c2] > 0 {
//					cntDict[c2]--
//					if cntDict[c2] == 0 {
//						bad--
//						spec--
//					}
//				}
//			}
//
//			if bad == 0 {
//				res = max395(res, right-left)
//			}
//		}
//	}
//
//	return res
//}

func max395(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
