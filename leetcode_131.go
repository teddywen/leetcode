/*
https://leetcode-cn.com/problems/palindrome-partitioning/
 */
package main

import "fmt"

func main() {
	var (
		s = "aab"
	)
	fmt.Printf("%+v", partition(s))
}

func partition(s string) [][]string {
	var (
		dict = make(map[string][][]string)
	)
	return palindromeDFS(s, dict)
}

func palindromeDFS(s string, memo map[string][][]string) [][]string {
	if len(s) == 0 {
		return nil
	}
	if v, ok := memo[s]; ok {
		return v
	}

	var (
		result [][]string
	)

	for i:=0; i<len(s); i++ {
		prefix, remain := s[:i+1], s[i+1:]
		if !isPalindrome(prefix) {
			continue
		}
		if len(remain) == 0 {
			result = append(result, []string{prefix})
			continue
		}
		groups := palindromeDFS(remain, memo)
		for _, group := range groups {
			combine := []string{prefix}
			combine = append(combine, group...)
			result = append(result, combine)
		}
	}

	memo[s] = result
	return result
}

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}
	var (
		left, right = 0, len(s)-1
	)
	for left <= right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
