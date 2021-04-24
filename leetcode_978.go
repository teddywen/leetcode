/*
https://leetcode-cn.com/problems/longest-turbulent-subarray/
*/
package main

import "fmt"

func main() {
	var (
		arr = []int{9, 4, 2, 10, 7, 8, 8, 1, 9}
	)
	fmt.Printf("%d", maxTurbulenceSize(arr))
}

// 滑动窗口解法
func maxTurbulenceSize(arr []int) int {
	var (
		res, i, j int
	)

	for j < len(arr) {
		func() {
			// 滑动右指针
			defer func() {
				j++
			}()

			if j-i == 0 {
				// 单元素也算长度为1的湍流子数组
				if res < j-i+1 {
					res = j - i + 1
				}
				return
			} else if j-i == 1 {
				// 双元素不相等才符合湍流子数组
				if arr[i] != arr[j] {
					if res < j-i+1 {
						res = j - i + 1
					}
					return
				}
			} else {
				// 符合湍流特性
				if (arr[j]-arr[j-1])*(arr[j-2]-arr[j-1]) > 0 {
					if res < j-i+1 {
						res = j - i + 1
					}
					return
				}
			}

			// 不符合湍流特性，则收缩
			i = j - 1
		}()
	}

	return res
}
