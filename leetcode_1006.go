/*
https://leetcode-cn.com/problems/clumsy-factorial/
*/
package main

import "fmt"

func main() {
	var (
		N = 4
	)
	fmt.Printf("%#v", clumsy(N))
}

// 迭代
func clumsy(N int) (ans int) {
	var (
		ops = []int{0, 1} // 0:'+' 1:'-'
		i   = 0
		n = N
	)
	for n > 0 {
		top := n == N
		if i == 0 {
			x := maxSafe(n)
			n--
			y := maxSafe(n)
			n--
			z := maxSafe(n)
			n--
			val := x * y / z
			if !top {
				val *= -1
			}
			ans += val
		} else {
			ans += n
			n--
		}
		i = (i + 1) % len(ops)
	}
	return
}

// 递归
//func clumsy(N int) (ans int) {
//	var (
//		ops = []int{0, 1} // 0:'+' 1:'-'
//	)
//	return _clumsyDFS(N, ops, 0, true)
//}
//
//func _clumsyDFS(n int, ops []int, i int, top bool) int {
//
//	if n < 1 {
//		return 0
//	}
//
//	switch ops[i] {
//	case 0:
//		x := maxSafe(n)
//		n--
//		y := maxSafe(n)
//		n--
//		z := maxSafe(n)
//		n--
//		val := x * y / z
//		if !top {
//			val *= -1
//		}
//		return val + _clumsyDFS(n, ops, (i+1)%len(ops), false)
//	case 1:
//		return n + _clumsyDFS(n-1, ops, (i+1)%len(ops), false)
//	default:
//		return 0
//	}
//}

func maxSafe(n int) int {
	if n < 1 {
		return 1
	} else {
		return n
	}
}
