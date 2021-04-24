/*
https://leetcode-cn.com/problems/number-of-islands/
*/
package main

import (
	"fmt"
)

func main() {
	var (
		grid = [][]byte{
			{'1', '1', '0', '0', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '1', '0', '0'},
			{'0', '0', '0', '1', '1'},
		}
	)
	fmt.Printf("%d", numIslands(grid))
}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	var (
		M, N     = len(grid), len(grid[0])
		uf       UF
		waterNum int
	)
	uf = NewUFPathCompressQuickUnion(M * N)

	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			if grid[i][j] == '1' {
				// 上
				if moveI := i - 1; moveI >= 0 && grid[moveI][j] == '1' {
					uf.Union(i*N+j, moveI*N+j)
				}
				// 下
				if moveI := i + 1; moveI < M && grid[moveI][j] == '1' {
					uf.Union(i*N+j, moveI*N+j)
				}
				// 左
				if moveJ := j - 1; moveJ >= 0 && grid[i][moveJ] == '1' {
					uf.Union(i*N+j, i*N+moveJ)
				}
				// 右
				if moveJ := j + 1; moveJ < N && grid[i][moveJ] == '1' {
					uf.Union(i*N+j, i*N+moveJ)
				}
			} else {
				waterNum++
			}
		}
	}
	return uf.Count() - waterNum
}

type UF interface {
	Union(p, q int)          // 合并 p 和 q
	Find(p int) int          // 查询 p 的集合标识
	Connected(p, q int) bool // 判断 p 和 q 是否在同一个集合内
	Count() int              // 共多少个集合
}

/////////////////////////////////////////////
//              quick-union impl           //
/////////////////////////////////////////////
type QuickUnion struct {
	roots []int
	count int
}

func NewUFQuickUnion(size int) *QuickUnion {
	uf := &QuickUnion{count: size, roots: make([]int, size)}
	// [0]=>0, [1]=>1, ..., [N]=>N
	for i := 0; i < size; i++ {
		uf.roots[i] = i
	}
	return uf
}

func (u *QuickUnion) Union(p, q int) {
	var (
		rootP = u.Find(p)
		rootQ = u.Find(q)
	)

	if rootP == rootQ {
		return
	}

	u.roots[rootP] = rootQ
	u.count--
}

func (u *QuickUnion) Find(p int) int {
	for u.roots[p] != p {
		p = u.roots[p]
	}
	return p
}

func (u *QuickUnion) Connected(p, q int) bool {
	return u.Find(p) == u.Find(q)
}

func (u *QuickUnion) Count() int {
	return u.count
}

/////////////////////////////////////////////
//         weighted quick-union impl       //
/////////////////////////////////////////////
type WeightedQuickUnion struct {
	roots []int
	sizes []int
	count int
}

func NewUFWeightedQuickUnion(size int) *WeightedQuickUnion {
	uf := &WeightedQuickUnion{count: size, roots: make([]int, size), sizes: make([]int, size)}
	// [0]=>0, [1]=>1, ..., [N]=>N
	for i := 0; i < size; i++ {
		uf.roots[i] = i
		uf.sizes[i] = 1
	}
	return uf
}

func (u *WeightedQuickUnion) Union(p, q int) {
	var (
		rootP = u.Find(p)
		rootQ = u.Find(q)
	)

	if rootP == rootQ {
		return
	}

	if u.sizes[rootP] < u.sizes[rootQ] {
		u.roots[rootP] = rootQ
		u.sizes[rootQ] += u.sizes[rootP]
	} else {
		u.roots[rootQ] = rootP
		u.sizes[rootP] += u.sizes[rootQ]
	}

	u.count--
}

func (u *WeightedQuickUnion) Find(p int) int {
	for u.roots[p] != p {
		p = u.roots[p]
	}
	return p
}

func (u *WeightedQuickUnion) Connected(p, q int) bool {
	return u.Find(p) == u.Find(q)
}

func (u *WeightedQuickUnion) Count() int {
	return u.count
}


/////////////////////////////////////////////
// weighted quick-union path compress impl //
/////////////////////////////////////////////
type PathCompressQuickUnion struct {
	roots []int
	sizes []int
	count int
}

func NewUFPathCompressQuickUnion(size int) *PathCompressQuickUnion {
	uf := &PathCompressQuickUnion{count: size, roots: make([]int, size), sizes: make([]int, size)}
	// [0]=>0, [1]=>1, ..., [N]=>N
	for i := 0; i < size; i++ {
		uf.roots[i] = i
		uf.sizes[i] = 1
	}
	return uf
}

func (u *PathCompressQuickUnion) Union(p, q int) {
	var (
		rootP = u.Find(p)
		rootQ = u.Find(q)
	)

	if rootP == rootQ {
		return
	}

	if u.sizes[rootP] < u.sizes[rootQ] {
		u.roots[rootP] = rootQ
		u.sizes[rootQ] += u.sizes[rootP]
	} else {
		u.roots[rootQ] = rootP
		u.sizes[rootP] += u.sizes[rootQ]
	}

	u.count--
}

func (u *PathCompressQuickUnion) Find(p int) int {
	var root = p
	for u.roots[root] != root {
		root = u.roots[root]
	}
	for u.roots[p] != p {
		u.roots[p], p = root, u.roots[p]
	}
	return p
}

func (u *PathCompressQuickUnion) Connected(p, q int) bool {
	return u.Find(p) == u.Find(q)
}

func (u *PathCompressQuickUnion) Count() int {
	return u.count
}