/*
https://leetcode-cn.com/problems/couples-holding-hands/
*/
package main

import "fmt"

func main() {
	var (
		row = []int{0, 2, 1, 3}
	)
	fmt.Printf("%d", minSwapsCouples(row))
}

var roots, sizes []int

func InitUF(len int) {
	roots = make([]int, len)
	for i := range roots {
		roots[i] = i
	}
	sizes = make([]int, len)
	for i := range sizes {
		sizes[i] = 1
	}
}

func FindUF(p int) int {
	root := p
	for roots[root] != root {
		root = roots[root]
	}
	for roots[p] != p {
		p, roots[p] = roots[p], root
	}
	return p
}

func UnionUF(x, y int) {
	var (
		rootX = FindUF(x)
		rootY = FindUF(y)
	)
	if rootX == rootY {
		return
	}

	if sizes[rootX] < sizes[rootY] {
		roots[rootX] = rootY
		sizes[rootY] += sizes[rootX]
	} else {
		roots[rootY] = rootX
		sizes[rootX] += sizes[rootY]
	}
}

// 并查集
func minSwapsCouples(row []int) int {
	var (
		counter = make([]int, len(row)/2)
		res     int
	)

	// 初始化并查集
	InitUF(len(row) / 2)

	// 并查集Union
	for i := 0; i < len(row); i += 2 {
		UnionUF(row[i]/2, row[i+1]/2)
	}

	// 统计连通分量的节点数
	for i := 0; i < len(row)/2; i++ {
		counter[FindUF(i)]++
	}

	for i := 0; i < len(counter); i++ {
		if counter[i] > 1 {
			res += counter[i] - 1
		}
	}
	return res
}

//// DFS
//func minSwapsCouples(row []int) int {
//	var (
//		visited = make([]int, len(row))
//		pos = make([]int, len(row)) // 情侣的位置
//		res int
//	)
//
//	// 初始化情侣的位置 情侣->座位
//	for i, v := range row {
//		pos[v] = i
//	}
//
//	for _, v := range row {
//		steps := minSwapsCouplesDFS(v, row, visited, pos)
//		if steps > 0 {
//			res += steps
//		}
//	}
//
//	return res
//}
//
//func minSwapsCouplesDFS(v int, row []int, visited []int, pos []int) int {
//	if visited[v] > 0 {
//		return -1 // -1 代表已成环，上一步不应该走到这里，因此返回-1来实现不记步
//	}
//
//	v2 := row[pos[v] ^ 1]
//	visited[v]++
//	visited[v2]++
//
//	return minSwapsCouplesDFS(v2 ^ 1, row, visited, pos) + 1
//}
