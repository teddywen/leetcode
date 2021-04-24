package explore

import "fmt"

// 给定一个数组，一直数组中只有一组重复元素，其余元素不重复。
// 找出改组重复元素
func main() {
	var (
		nums = []int{1, 2, 3, 4, 5, 3}
		res  int
	)
	for _, n := range nums {
		res ^= n
	}
	fmt.Println(res)
}

// 排序+快慢指针
// 时间复杂度: O(nlogn + n)
// 空间复杂度: O(1)

// 快排三向切分
// 时间复杂度: O(nlogn)
// 空间复杂度: O(1)

// 哈希计数
// 时间复杂度: O(n)
// 空间复杂度: O(n)

// ?
// 时间复杂度: O(n)
// 空间复杂度: O(1)
