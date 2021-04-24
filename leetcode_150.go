/*
https://leetcode-cn.com/problems/evaluate-reverse-polish-notation/
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var (
		tokens = []string{"4","13","5","/","+"}
	)
	fmt.Printf("%#v", evalRPN(tokens))
}

func evalRPN(tokens []string) int {
	if len(tokens) == 0 {
		return 0
	}
	var (
		stack []string
	)
	for _, v := range tokens {
		if v == "+" || v == "-" || v == "*" || v == "/" {
			num1, _ := strconv.Atoi(stack[len(stack)-2])
			num2, _ := strconv.Atoi(stack[len(stack)-1])
			num := 0
			if v == "+" {
				num = num1 + num2
			} else if v == "-" {
				num = num1 - num2
			} else if v == "*" {
				num = num1 * num2
			} else {
				num = num1 / num2
			}
			stack[len(stack)-2] = strconv.Itoa(num)
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, v)
		}
	}
	result, _ := strconv.Atoi(stack[0])
	return result
}
