package greedy_algorithm

import (
	"fmt"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	nums, k := []int{1, 1, 1, 2, 2, 2, 3}, 2
	res := topKFrequent(nums, k)
	fmt.Println(res)
}

func topKFrequent(nums []int, k int) []int {
	// 确定每个数出现的频率
	m := make(map[int]int)
	var res []int
	for _, v := range nums {
		m[v]++
	}
	// map[1:3 2:3 3:1]
	fmt.Println(m)

	// 创建桶，由于k不会大于nums长度，所以桶的大小为len(nums)+1
	buckets := make([][]int, len(nums)+1)
	for k, v := range m {
		buckets[v] = append(buckets[v], k)
	}
	// [[] [3] [] [1 2] [] [] [] []]
	fmt.Println(buckets, len(buckets))

	// 从最大桶开始取数据
	for i := len(buckets) - 1; k > 0; i-- {
		for _, v := range buckets[i] {
			res = append(res, v)
			k--
		}
	}

	return res
}
