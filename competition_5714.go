/*
https://leetcode-cn.com/problems/evaluate-the-bracket-pairs-of-a-string/
*/
package main

import (
	"fmt"
)

func main() {
	var (
		s         = "(a)(b)"
		knowledge = [][]string{{"a", "b"}, {"b", "a"}}
	)
	fmt.Printf("%s", evaluate(s, knowledge))
}

func evaluate(s string, knowledge [][]string) string {
	var (
		stack []byte
		dict  = make(map[string]string)
	)
	for _, v := range knowledge {
		dict[v[0]] = v[1]
	}
	for i := range s {
		if s[i] == ')' {
			var key []byte
			for {
				c := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if c == '(' {
					k, j := 0, len(key)-1
					for k <= j {
						key[k], key[j] = key[j], key[k]
						k++
						j--
					}
					keyStr := string(key)
					var val []byte
					if v, ok := dict[keyStr]; ok {
						val = []byte(v)
					} else {
						val = []byte{'?'}
					}
					stack = append(stack, val...)
					break
				} else {
					key = append(key, c)
				}
			}
		} else {
			stack = append(stack, s[i])
		}
	}

	return string(stack)
}
