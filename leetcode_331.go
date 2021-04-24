/*
https://leetcode-cn.com/problems/verify-preorder-serialization-of-a-binary-tree/
*/
package main

import (
	"fmt"
)

func main() {
	var (
		preorder = "1,#,#,1"
	)
	fmt.Printf("%t", isValidSerialization(preorder))
}

// simulate preorder dfs
func isValidSerialization(preorder string) bool {
	var (
		ptr *int
	)
	i := 0
	ptr = &i

	return preorderDFS(preorder, ptr) && *ptr == len(preorder)
}

func preorderDFS(preorder string, ptr *int) bool {
	if *ptr >= len(preorder) {
		return false
	}
	if preorderPop(preorder, ptr) {
		return true
	}
	return preorderDFS(preorder, ptr) && preorderDFS(preorder, ptr)
}

func preorderPop(preorder string, ptr *int) bool {
	var isEnd bool
	for ; *ptr < len(preorder); *ptr++ {
		if preorder[(*ptr)] == ',' {
			*ptr++
			break
		}
		if preorder[(*ptr)] == '#' {
			isEnd = true
		}
	}
	return isEnd
}
