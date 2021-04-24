/*
https://leetcode-cn.com/problems/number-of-valid-words-for-each-puzzle/
*/
package main

import "fmt"

func main() {
	var (
		words   = []string{"aaaa", "asas", "able", "ability", "actt", "actor", "access"}
		puzzles = []string{"a", "abrodyz", "abslute", "absoryz", "actresz", "gaswxyz"}
	)
	fmt.Printf("%#v", findNumOfValidWords(words, puzzles))
}

// 遍历比特位子集法
func findNumOfValidWords(words []string, puzzles []string) []int {
	var (
		wordsCnt = make(map[int]int)
		result   = make([]int, len(puzzles))
	)
	for _, word := range words {
		var bitSet int
		for _, c := range word {
			bitSet |= 1 << (c-'a')
		}
		wordsCnt[bitSet]++
	}
	for i, puzzle := range puzzles {
		var bitSet, firstBit, mask, cnt int
		for j, c := range puzzle {
			if j > 0 {
				bitSet |= 1 << (c-'a')
			} else {
				firstBit = 1 << (c-'a')
			}
		}
		mask = bitSet
		for cnt += wordsCnt[bitSet|firstBit]; bitSet > 0; {
			bitSet = (bitSet - 1) & mask
			cnt += wordsCnt[bitSet|firstBit]
		}
		result[i] = cnt
	}
	return result
}

// O(m*n)
//func findNumOfValidWords(words []string, puzzles []string) []int {
//	var (
//		bitDict    = make([]int, 26)
//		wordsBit   = make([]int, len(words))
//		puzzlesBit = make([][2]int, len(puzzles))
//		result     = make([]int, len(puzzles))
//	)
//	for i := range bitDict {
//		bitDict[i] = 1 << i
//	}
//
//	for i := range wordsBit {
//		var bit int
//		for _, c := range words[i] {
//			bit |= bitDict[c-'a']
//		}
//		wordsBit[i] = bit
//	}
//	for i := range puzzlesBit {
//		var bit, capBit int
//		for j, c := range puzzles[i] {
//			if j == 0 {
//				capBit = bitDict[c-'a']
//			}
//			bit |= bitDict[c-'a']
//		}
//		puzzlesBit[i][0], puzzlesBit[i][1] = bit, capBit
//	}
//
//	for i := range puzzlesBit {
//		var count int
//		for j := range wordsBit {
//			if (puzzlesBit[i][0] & wordsBit[j]) != wordsBit[j] {
//				continue
//			}
//			if (puzzlesBit[i][1] & wordsBit[j]) == 0 {
//				continue
//			}
//			count++
//		}
//		result[i] = count
//	}
//
//	return result
//}
