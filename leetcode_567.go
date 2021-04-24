/*
https://leetcode-cn.com/problems/permutation-in-string/
*/
package main

import "fmt"

func main() {
	var (
		s1 = "hello"
		s2 = "ooolleoooleh"
	)
	fmt.Printf("%t", checkInclusion(s1, s2))
}

func checkInclusion(s1 string, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}

	var (
		// cnt1: s1一共多少种字符
		// cnt2: s2数量达标的字符有多少个
		// 当 len(s2[i:j]) == len(s1) && cnt1 == cnt2 时，代表 s2 包含 s1 的排列。
		cnt1, cnt2  int
		left, right int
		num1        = [26]int{} // s1每个字符的数量统计
		num2        = [26]int{} // s2每个字符的数量统计
	)

	// 根据定义初始化 num1、cnt1
	for _, c := range s1 {
		if num1[c-'a'] == 0 {
			cnt1++
		}
		num1[c-'a']++
	}

	for right < len(s2) {
		// 右滑
		c := s2[right]
		right++

		// 更新 num2 和 cnt2
		if num1[c-'a'] > 0 {
			num2[c-'a']++
			if num2[c-'a'] == num1[c-'a'] {
				cnt2++
			}
		}

		for right-left > len(s1) {
			// 收缩
			c2 := s2[left]
			left++

			// 更新 num2 和 cnt2
			if num1[c2-'a'] > 0 {
				if num2[c2-'a'] > 0 {
					if num2[c2-'a'] == num1[c2-'a'] {
						cnt2--
					}
					num2[c2-'a']--
				}
			}
		}

		if right-left == len(s1) && cnt1 == cnt2 {
			return true
		}
	}

	return false
}
