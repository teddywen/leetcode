/*
https://leetcode-cn.com/problems/basic-calculator-ii/
*/
package main

import (
	"fmt"
)

func main() {
	var (
		s = "0"
	)
	fmt.Printf("%d", calculate(s))
}

// 讨巧正负号压栈
func calculate(s string) int {
	var (
		stack []int
		sign  = 1
		result int
	)

	for i := 0; i < len(s); {
		switch s[i] {
		case ' ':
			i++
		case '+':
			sign = 1
			i++
		case '-':
			sign = -1
			i++
		case '*':
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			i++
			var num int
			for ; i < len(s) && ((s[i] >= '0' && s[i] <= '9') || s[i] == ' '); i++ {
				if s[i] == ' ' {
					continue
				}
				num = num*10 + int(s[i]-'0')
			}
			stack = append(stack, last*num)
		case '/':
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			i++
			var num int
			for ; i < len(s) && ((s[i] >= '0' && s[i] <= '9') || s[i] == ' '); i++ {
				if s[i] == ' ' {
					continue
				}
				num = num*10 + int(s[i]-'0')
			}
			stack = append(stack, last/num)
		default:
			var num int
			for ; i < len(s) && s[i] >= '0' && s[i] <= '9'; i++ {
				num = num*10 + int(s[i]-'0')
			}
			stack = append(stack, sign*num)
			sign = 1
		}
	}

	for _, v := range stack {
		result += v
	}
	return result
}
