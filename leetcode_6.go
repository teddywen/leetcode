/*
https://leetcode-cn.com/problems/zigzag-conversion/
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		s       = "AB"
		numRows = 1
	)
	fmt.Printf("%s", convert(s, numRows))
}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	var (
		sb     = strings.Builder{}
		binary, curRow = 1, 0
		matrix = make([][]byte, numRows)
	)
	for i := range s {
		matrix[curRow] = append(matrix[curRow], s[i])
		curRow += binary
		if curRow == 0 || curRow == numRows-1 {
			binary = ^binary + 1
		}
	}
	for i := range matrix {
		for j := range matrix[i] {
			sb.WriteByte(matrix[i][j])
		}
	}
	return sb.String()
}
