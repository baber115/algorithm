package greedy_algorithm

import (
	"container/heap"
	"fmt"
	"testing"
)

// https://leetcode.cn/problems/maximum-average-pass-ratio/solutions/2118606/zui-da-ping-jun-tong-guo-lu-by-leetcode-dm7y3/

/*
优先队列
难点：公式推导
*/
func TestMaxAverageRatio(t *testing.T) {
	classes, extraStudents := [][]int{{1, 2}, {3, 5}, {2, 2}}, 2
	res := maxAverageRatio(classes, extraStudents)
	fmt.Println(res)
}

func maxAverageRatio(classes [][]int, extraStudents int) (ans float64) {
	h := hp(classes)
	heap.Init(&h)
	for ; extraStudents > 0; extraStudents-- {
		h[0][0]++
		h[0][1]++
		heap.Fix(&h, 0)
	}

	for _, v := range h {
		ans += float64(v[0]) / float64(v[1])
	}

	return ans / float64(len(classes))
}

type hp [][]int

func (h hp) Len() int { return len(h) }
func (h hp) Less(i, j int) bool {
	a, b := h[i], h[j]
	return (a[1]-a[0])*b[1]*(b[1]+1) > (b[1]-b[0])*a[1]*(a[1]+1)
}
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (hp) Push(interface{})     {}
func (hp) Pop() (_ interface{}) { return }
