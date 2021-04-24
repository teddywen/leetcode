/*
https://leetcode-cn.com/problems/flipping-an-image/
*/
package main

import "fmt"

func main() {
	var (
		nums = [][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}
	)
	fmt.Printf("%#v", flipAndInvertImage(nums))
}

func flipAndInvertImage(A [][]int) [][]int {
	for _, row := range A {
		i, j := 0, len(row)-1
		for i <= j {
			row[i], row[j] = row[j]^1, row[i]^1
			i++
			j--
		}
	}
	return A
}

// 官方写法
//func flipAndInvertImage(A [][]int) [][]int {
//	for _, row := range A {
//		i, j := 0, len(row)-1
//		for i < j {
//			// 细节: 如果左右不相等则没必要操作，因为翻转加反转正好等于原值。
//			if row[i] == row[j] {
//				row[i], row[j] = row[i]^1, row[j]^1
//			}
//			i++
//			j--
//		}
//		// 中间元素不能漏
//		if i == j {
//			row[i] ^= 1
//		}
//	}
//	return A
//}
