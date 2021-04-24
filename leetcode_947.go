/*
https://leetcode-cn.com/problems/most-stones-removed-with-same-row-or-column/
*/
package main

import "fmt"

func main() {
	var (
		stones = [][]int{{0, 1}, {1, 1}}
	)
	fmt.Printf("%d", removeStones(stones))
}

func removeStones(stones [][]int) int {
	var rows int
	for _, pt := range stones {
		rows = max947(rows, pt[0]+1)
	}

	uf := NewUF947()
	for _, pt := range stones {
		uf.Union(pt[0], rows+pt[1])
	}

	return len(stones) - uf.Count()
}

func max947(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

type UF947 struct {
	roots map[int]int
	sizes map[int]int
	count int
}

func NewUF947() *UF947 {
	return &UF947{roots: make(map[int]int), sizes: make(map[int]int)}
}

func (uf *UF947) Find(x int) int {
	var (
		root = x
	)
	if _, ok := uf.roots[x]; !ok {
		uf.roots[x] = x
		uf.sizes[x] = 1
		uf.count++
	}
	for uf.roots[root] != root {
		root = uf.roots[root]
	}
	for uf.roots[x] != x {
		uf.roots[x], x = root, uf.roots[x]
	}
	return x
}

func (uf *UF947) Union(x, y int) {
	xRoot := uf.Find(x)
	yRoot := uf.Find(y)
	if xRoot == yRoot {
		return
	}
	if uf.sizes[xRoot] > uf.sizes[yRoot] {
		uf.roots[yRoot] = xRoot
		uf.sizes[xRoot] += uf.sizes[yRoot]
	} else {
		uf.roots[xRoot] = yRoot
		uf.sizes[yRoot] += uf.sizes[xRoot]
	}
	uf.count--
}

func (uf *UF947) Count() int {
	return uf.count
}
