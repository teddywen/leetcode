/*
https://leetcode-cn.com/problems/reverse-bits/
*/
package main

import "fmt"

func main() {
	var (
		num = uint32(0b00000010100101000001111010011100)
	)
	fmt.Printf("%d", reverseBits(num))
}

func reverseBits(num uint32) uint32 {
	var (
		i, j = 0, 31
	)
	for i <= j {
		num = reverseTwoBits(num, i, j)
		i++
		j--
	}
	return num
}
func reverseTwoBits(num uint32, i, j int) uint32 {
	iBitZero := (num & (1 << i)) == 0
	jBitZero := (num & (1 << j)) == 0
	if iBitZero {
		num &= ^(1 << j)
	} else {
		num |= 1 << j
	}
	if jBitZero {
		num &= ^(1 << i)
	} else {
		num |= 1 << i
	}
	return num
}
