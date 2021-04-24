/*
https://leetcode-cn.com/problems/basic-calculator/
*/
package main

import (
	"fmt"
)

func main() {
	var (
		s = "(1+(4+5 +2)-3) +(6+8)"
	)
	fmt.Printf("%#v", calculate(s))
}

// 讨巧正负号压栈
func calculate(s string) int {
	var (
		stack  = []int{1}
		sign   = 1 // 当前符号
		i, res int
	)
	for i < len(s) {
		c := s[i]
		switch c {
		case ' ': // do nothing
			i++
		case '(':
			stack = append(stack, sign * stack[len(stack)-1])
			sign = 1
			i++
		case ')':
			stack = stack[:len(stack)-1]
			i++
		case '+':
			sign = 1
			i++
		case '-':
			sign = -1
			i++
		default:
			num := 0
			for ; i < len(s) && s[i] >= '0' && s[i] <= '9'; i++ {
				num = num*10 + int(s[i]-'0')
			}
			res += sign * stack[len(stack)-1] * num
		}
	}
	return res
}

// 常规压栈计算
//func calculate(s string) int {
//	var (
//		stack = make([]byte, 0)
//		memo  = make([]int, 0)
//	)
//
//	stack = append(stack, '(')
//
//	for i := range s {
//		if isBlank(s[i]) {
//			continue
//		}
//		stack = append(stack, s[i])
//		if isCloseBrace(s[i]) {
//			calculateBrace(&stack, &memo)
//		}
//	}
//
//	stack = append(stack, ')')
//	calculateBrace(&stack, &memo)
//
//	return memo[0]
//}
//
//func calculateBrace(stack *[]byte, memo *[]int) {
//	// remove close brace
//	*stack = (*stack)[:len(*stack)-1]
//
//	var sum int
//	for !isOpenBrace((*stack)[len(*stack)-1]) {
//		num, numBit := 0, 0
//
//		// 提取数字
//		for isNumeric((*stack)[len(*stack)-1]) {
//			n := toNumeric((*stack)[len(*stack)-1])
//			num += n * int(math.Pow(float64(10), float64(numBit)))
//			numBit++
//			*stack = (*stack)[:len(*stack)-1]
//		}
//		// 提取寄存数
//		if isMemo((*stack)[len(*stack)-1]) {
//			num = (*memo)[len(*memo)-1]
//			*memo = (*memo)[:len(*memo)-1]
//			*stack = (*stack)[:len(*stack)-1]
//		}
//		// 提取符号
//		if len(*stack) > 0 && isSymbol((*stack)[len(*stack)-1]) {
//			if isMinus((*stack)[len(*stack)-1]) {
//				num *= -1
//			}
//			*stack = (*stack)[:len(*stack)-1]
//		}
//		sum += num
//	}
//
//	// remove open brace
//	*stack = (*stack)[:len(*stack)-1]
//
//	*memo = append(*memo, sum)
//	*stack = append(*stack, '$')
//}
//
//func isMemo(c byte) bool {
//	return c == '$'
//}
//
//func isOpenBrace(c byte) bool {
//	return c == '('
//}
//
//func isCloseBrace(c byte) bool {
//	return c == ')'
//}
//
//func isNumeric(c byte) bool {
//	return c >= '0' && c <= '9'
//}
//
//func toNumeric(c byte) int {
//	return int(c - '0')
//}
//
//func isMinus(c byte) bool {
//	return c == '-'
//}
//
//func isPlus(c byte) bool {
//	return c == '+'
//}
//
//func isSymbol(c byte) bool {
//	return isMinus(c) || isPlus(c)
//}
//
//func isBlank(c byte) bool {
//	return c == ' '
//}
