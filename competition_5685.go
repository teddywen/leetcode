/*
https://leetcode-cn.com/contest/weekly-contest-229/problems/merge-strings-alternately/
 */
package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		word1 = "abcd"
		word2 = "pq"
	)
	fmt.Printf("%#v", mergeAlternately(word1, word2))
}

func mergeAlternately(word1 string, word2 string) string {
	var (
		sb = strings.Builder{}
		i, j int
	)

	for i < len(word1) || j < len(word2) {
		if i < len(word1) {
			sb.WriteByte(word1[i])
		}
		if j < len(word2) {
			sb.WriteByte(word2[j])
		}
		i++
		j++
	}

	return sb.String()
}
